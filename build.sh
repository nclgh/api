RUN_NAME="bear.server.like"
KITE_PATH="code.byted.org/kite"
REPO_PATH="code.byted.org/server/like"

mkdir -p output/bin output/conf
cp script/bootstrap.sh output
cp script/bootstrap_staging.sh output
chmod +x output/bootstrap.sh
chmod +x output/bootstrap_staging.sh
cp -r conf output/
find conf/ -type f ! -name "*_local.*" | xargs -I{} cp {} output/conf/

export GO15VENDOREXPERIMENT="1"

GIT_SHA=$(git rev-parse --short HEAD || echo "GitNotFound")
DATE=$(date '+%Y%m%d%H%M%S')
VERSION=${GIT_SHA}-${DATE}

LINK_OPERATOR="="

go build -ldflags "-X ${REPO_PATH}/vendor/${KITE_PATH}/kite.ServiceVersion${LINK_OPERATOR}${VERSION}" -o output/bin/${RUN_NAME}