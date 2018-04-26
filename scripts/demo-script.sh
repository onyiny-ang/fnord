#!/bin/bash
#==========================================================
#  Script name -- demo-script.sh --
# DESCRIPTION:  Federation v2 Demo
#
#
#  Author:  Lindsey Tulloch , ltulloch@redhat.com
#  CREATED:  2018-04-25 03:42:58 PM EDT

set -e

CONTEXT=us-west
CONTEXT2=us-east
NS=test-namespace

printf "\nkubectl -n federation describe federatedclusters\n"
kubectl -n federation describe federatedclusters

read -n 1 -s

printf "\nkubectl create -f ./configs/\n"
kubectl create -f ./configs/

read -n 1 -s

printf "Namespaces for existing contexts\n"
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} get ns test-namespace; echo; echo; done
read -n 1 -s
printf "Configmaps for existing contexts\n"
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} get configmaps; echo; echo; done
printf "Secrets for existing contexts\n"
read -n 1 -s
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} get secrets; echo; echo; done
printf "Replicasets for existing contexts\n"
read -n 1 -s
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} get rs; echo; echo; done
read -n 1 -s

kubectl -n ${NS} edit federatednamespaceplacement ${NS}

sleep 7s

printf " Replicasets showing deleted us-east\n"
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} -n ${NS} get rs; echo; echo; done

read -n 1 -s

kubectl -n ${NS} edit federatednamespaceplacement ${NS}
sleep 7s

printf "Replicasets showing us-east restored\n"
for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} -n ${NS} get rs; echo; echo; done

sleep 7s

for i in ${CONTEXT} ${CONTEXT2}; do echo; echo ------------ ${i} ------------; echo; kubectl --context ${i} -n ${NS} get rs; echo; echo; done
