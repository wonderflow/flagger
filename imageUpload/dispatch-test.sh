regions=(cn-shanghai cn-beijing cn-hangzhou cn-qingdao)
for region in ${regions[@]}
do
	echo ${region}
	docker tag weaveworks/flagger:${TAG} registry.${region}.aliyuncs.com/${NAMESPACE}/${REPO}
	docker push registry.${region}.aliyuncs.com/${NAMESPACE}/${REPO}
done