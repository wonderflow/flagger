regions=(ap-southeast-1 ap-northeast-1 cn-beijing cn-hangzhou cn-hongkong cn-qingdao cn-shanghai cn-shenzhen cn-zhangjiakou cn-chengdu cn-huhehaote)
for region in ${regions[@]}
do
	echo ${region}
	docker tag weaveworks/flagger:${TAG} registry.${region}.aliyuncs.com/${NAMESPACE}/${REPO}
	docker push registry.${region}.aliyuncs.com/${NAMESPACE}/${REPO}
done