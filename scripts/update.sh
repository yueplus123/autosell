#!/usr/bin/env bash
# 一键更新服务器上的程序到 GitHub 最新发布版本。
#
# 用法（在服务器上以 root 执行）：
#   curl -fsSL https://raw.githubusercontent.com/yueplus123/autosell/main/scripts/update.sh | bash
#
# 可选环境变量：
#   VERSION=v1.1.2      指定版本（默认最新发布）
#   INSTALL_DIR=/path   程序目录（默认 /opt/dujiao-next）
#   SERVICE=name        systemd 服务名（默认 dujiao-next）
#
# 做的事：下载 → 校验 → 备份当前程序与数据目录 → 替换 → 重启 → 健康检查（失败自动回滚）
#        → 初始化业务分类（可重复执行，已存在的分类不改动）。
set -euo pipefail

REPO="yueplus123/autosell"
INSTALL_DIR="${INSTALL_DIR:-/opt/dujiao-next}"
SERVICE="${SERVICE:-dujiao-next}"
BINARY="${INSTALL_DIR}/dujiao-next"
BACKUP_DIR="/root/dujiao-backups"

info() { printf '\033[1;32m==>\033[0m %s\n' "$*"; }
warn() { printf '\033[1;33m[注意]\033[0m %s\n' "$*" >&2; }
die() { printf '\033[1;31m[失败]\033[0m %s\n' "$*" >&2; exit 1; }

[ "$(id -u)" -eq 0 ] || die "请用 root 执行（或在命令前加 sudo）"
command -v curl >/dev/null || die "缺少 curl"
command -v tar >/dev/null || die "缺少 tar"

if [ ! -x "$BINARY" ] || ! systemctl cat "${SERVICE}.service" >/dev/null 2>&1; then
  if command -v docker >/dev/null 2>&1 && docker ps --format '{{.Image}} {{.Names}}' 2>/dev/null | grep -qi dujiao; then
    die "检测到程序以 Docker 方式运行，本脚本只适用于一键安装脚本（systemd）部署。"
  fi
  die "没有在 ${INSTALL_DIR} 找到程序或 ${SERVICE} 服务。如安装在别处，请用 INSTALL_DIR=/你的目录 重新执行。"
fi

case "$(uname -m)" in
  x86_64 | amd64) ARCH="x86_64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) die "不支持的 CPU 架构：$(uname -m)" ;;
esac

# ---------- 确定版本 ----------
if [ -z "${VERSION:-}" ]; then
  info "查询最新版本…"
  VERSION="$(curl -fsSL "https://api.github.com/repos/${REPO}/releases/latest" \
    | grep -m1 '"tag_name"' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')" || true
  [ -n "${VERSION:-}" ] || die "查询最新版本失败（GitHub 上可能还没有发布，或服务器访问 GitHub 受限）"
fi
info "目标版本：${VERSION}"

# ---------- 下载与校验 ----------
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT
ASSET="dujiao-next_${VERSION}_Linux_${ARCH}.tar.gz"
BASE_URL="https://github.com/${REPO}/releases/download/${VERSION}"

info "下载 ${ASSET}…"
curl -fL --retry 3 -o "${WORK}/${ASSET}" "${BASE_URL}/${ASSET}" || die "下载失败：${BASE_URL}/${ASSET}"

SUMS="dujiao-next_${VERSION#v}_checksums.txt"
if curl -fsL --retry 3 -o "${WORK}/${SUMS}" "${BASE_URL}/${SUMS}"; then
  ( cd "$WORK" && grep " ${ASSET}\$" "$SUMS" | sha256sum -c - >/dev/null ) || die "文件校验失败，已中止（未做任何改动）"
  info "校验通过"
else
  warn "没有找到校验文件，跳过校验"
fi

mkdir -p "${WORK}/pkg"
tar xzf "${WORK}/${ASSET}" -C "${WORK}/pkg"
NEW_BIN="$(find "${WORK}/pkg" -type f -name dujiao-next | head -n1)"
[ -n "$NEW_BIN" ] || die "安装包里没有找到程序文件"
chmod 755 "$NEW_BIN"

# ---------- 备份 ----------
STAMP="$(date +%Y%m%d-%H%M%S)"
mkdir -p "$BACKUP_DIR"
info "停止服务并备份到 ${BACKUP_DIR}/…"
systemctl stop "$SERVICE"
cp -a "$BINARY" "${BACKUP_DIR}/dujiao-next.${STAMP}"
tar czf "${BACKUP_DIR}/data.${STAMP}.tar.gz" -C "$INSTALL_DIR" \
  --exclude='./dujiao-next' --exclude='./logs' --exclude='*.tar.gz' . 2>/dev/null \
  || warn "数据目录备份不完整，请留意"

rollback() {
  warn "新版本启动失败，正在回滚到更新前的版本…"
  systemctl stop "$SERVICE" || true
  cp -a "${BACKUP_DIR}/dujiao-next.${STAMP}" "$BINARY"
  systemctl start "$SERVICE" || true
  die "已回滚。可用 journalctl -u ${SERVICE} -n 100 查看失败原因。"
}

# ---------- 替换并启动 ----------
info "替换程序并启动…"
install -m 755 "$NEW_BIN" "$BINARY"
systemctl start "$SERVICE" || rollback
sleep 5
systemctl is-active --quiet "$SERVICE" || rollback
info "服务已启动"

# ---------- 初始化业务分类 ----------
info "初始化业务分类…"
( cd "$INSTALL_DIR" && "$BINARY" admin seed-categories ) || warn "分类初始化失败，可稍后手动执行：cd ${INSTALL_DIR} && ./dujiao-next admin seed-categories"

# 只保留最近 5 份备份
ls -1t "${BACKUP_DIR}"/dujiao-next.* 2>/dev/null | tail -n +6 | xargs -r rm -f
ls -1t "${BACKUP_DIR}"/data.*.tar.gz 2>/dev/null | tail -n +6 | xargs -r rm -f

info "更新完成：${VERSION}"
echo "   如需回滚：systemctl stop ${SERVICE} && cp ${BACKUP_DIR}/dujiao-next.${STAMP} ${BINARY} && systemctl start ${SERVICE}"
