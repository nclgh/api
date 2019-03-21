service=lakawei_api

rm -rf output
mkdir -p output/conf
cp -rf conf ./output/
export GO15VENDOREXPERIMENT="1"
go build
mv $service ./output/

# run
cd ./output
./$service