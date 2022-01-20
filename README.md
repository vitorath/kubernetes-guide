# Kubernetes

## Introdução

Kubernetes (K8s) é um produto Open Source, desenvolvido pela Google, utilizado para automatizar a implantação, o dimensionamento e o gerencimanto de aplicativos em contêiner. Logo o melhor cloud provider atualmente executar o Kubernetes é o GCP, pois é o serviço fornecido pela própria Google.

O Kuberntes é disponibilizados por meio de um conjunto de APIs, no qual, normalmente acessamos utilizando o CLI **kubectl**.Nele tudo é baseado em estados ao qual o desenvolvedor configura os estados de um cada objeto, em arquivo(s) de configuração(ões), e o kubernetes aplica todas as configurações criando assim o ambiente. Alguns dos objetos existentes no kubernetes são:

- Pod: é uma unidade que contém os container provisionados. Além disso, um pod é processo rodando dentro de um cluster
- ReplicaSet: é um objeto utilizado para criar multiplas replicas de um pod, auxiliando assim no principio de resiliencia de um serviço, isto é, caso um pod caia o outro ainda continuara funcionando e aplicação continuará funcionando;
- Deployment: tem o objetivo de provisionar os **pods**, para isso, é necessário informar para ele quantas replicas serão provisionadas (ReplicaSet).
- Services: é uma forma de agregar um conjunto de pods para então implantar políticas de visibilidade

O Kuberntes trabalha através de **clusters** que contém um conjunto de **nodes**, logo, ele tem um cluster denomidado **master** que controla os nodes, que contem os seguintes serviços:

- Kube-apiserver
- Kube-controller-manager
- Kube-scheduler

Além disso, existes outros nodes que não são o master, sendo eles:

- Kubelet
- Kubeproxy

Um node consiste em uma máquina e cada possui uma quantidade de vCPU e Memória para serem utilizada. Com estes recursos mensurados o Kubernetes consegue medir a quantidade utilizada e gerencia-los adequadamente de acordo com o que foi configurado. Logo, caso a aplicação tenha um node e todos o recursos foram consumidos ou não tenha suficiante o kubernetes informa que não tem recursos suficientes para aplicar a operação. Contudo, caso exista outro **node** o kubernetes utilizará este outro **node**. Atém disso, caso um **node** caia o kubernetes irá recriar todas as configurações que estavem neste **node** dentro de outro que tenha recursos disponíveis.

Uma boa prática é manter um **pod** por **node**, contudo, pode existir situações na qual um **node** tenha mais de um **pod**, embora sejam cenários bem específicos.

## Kind - uma ferramenta para praticar com Kubernetes

[Kind](https://kind.sigs.k8s.io/docs/user/quick-start) é uma ferramenta que executa localmente clusters de Kubernetes usando "nodes" containers Docker. Contudo, para ter acesso aos clusters que são gerandos pelo **kind** é mais conveniente utilizar o [kubectl](https://kubernetes.io/docs/tasks/tools/) que é client do kuberntes que se comunica com o servidor do kubernetes.

### Criando clusters com o Kind

Para criar um cluster com um único **node** utilizando o kind execute o comando:

```sh
kind create cluster
```

Contudo, existe outra forma de criar clusters no kind por meio de um arquivos de configuração, neste caso [k8s/kind-config.yaml](./k8s/kind-config.yaml)

```sh
kind create cluster --config=k8s/kind-config.yaml
```

As informações referente aos clusters que sua máquina está configurada para conectar-se são encontradas no arquivo **~/.kube/config**, lembrando que este arquivo irá existir somente se houver alguma conexão configurada. No casos estas informações podem ser acessadas por meio do **kubectl**, por exemplo para visualizar todos os clusters configurados

```sh
kubectl config get-clusters
```

E conseguir acessar um cluster

```sh
# Lembrando que "kind-kind" é o nome do contexto
kubectl cluster-info --context kind-kind
```

Outra forma para mudar de contexto

```sh
# Lembrando que "kind-fullcycle" é o nome do contexto
kubectl config use-context kind-fullcycle
```

Para verificamos se estamos conectados o comando abaixo deve retornar um ou mais nodes

```sh
kubectl get nodes
```

Contudo caso, sejá necessário deletar um cluster do kind devemos saber qual o nome do cluster (por padrão, caso não tenha especificado o nome é kind) e posteriormente executar o comando de delete

```sh
# Lista todos os clusters criados pelo kind
kind get clusters

# Deleta o cluster com o nome "kind"
kind delete clusters kind
```

**NOTA**: Para facilitar o gerenciamento das conexões do kubernetes pode utilizar a extensão do vscode (kubernetes-tools)[https://marketplace.visualstudio.com/items?itemName=ms-kubernetes-tools.vscode-kubernetes-tools]. **Lembrando, que esta ferramenta gerencia somente as conexões com o kubernetes e não o _kind_ em si**.

## Pod

**Pod** é um conjunto de um ou mais containers, sendo assim a menor unidade de uma aplicação kubernetes. Contudo o normal é executar somente um container por **pod**. Contudo, um **pod** por si só não será recriado caso ocorra algum erro e ele seja derrubado.

A forma mais comum de criar um **pod** é por meio de um arquivo de configuração, neste caso (k8s/kube-pod.yaml)[./k8s/kube-pod.yaml]. Para aplica-lo utilize o comando **apply** do **kubectl**.

```sh
# Lembrando que "k8s/kube-pod.yaml" é o caminho até o arquivo de configuração.
kubectl apply -f k8s/kube-pod.yaml
```

Após um **pod** ser criado é possível visualiza-lo(s) por meio do comando **get pods** ou **get po**. Caso queira deleta-lo utilize o comando delete

```sh
kubectl get pods
kubectl get po

# Lembrando que "goserver" é o nome do pod
kubectl delete pod goserver
```

**Note**: Para conseguir visualizar a execução da aplicação (não sendo o ideal) basta fazer um redirecionamento de portas entre a rede local e a do pod no kubernetes

```sh
# Sendo "/goserver" o nome da aplicação
# 8000 a porta da rede local (externa ao kubernetes)
# 80 a porta interna do pod no kubernetes
kubectl port-forward pod/goserver 8000:80
```

## ReplicaSet

**ReplicaSet** gerencia os **pods** mantendo em execução o número especificado de replicas. O arquivo (k8s/kube-replicaset)[./k8s/kube-replicaset/] é um exemplo de ReplicaSet e para executa-lo utilize o apply, similar ao pod.

```sh
# Sendo "k8s/kube-replicaset.yaml" o caminho do arquivo de configuração
kubectl apply -f k8s/kube-replicaset.yaml
```

Para visualizar o(s) ReplicaSets criado execute:

```sh
kubectl get replicasets
```

# Informações que podem ser utilizadas

Borg > Omega > Kubernetes (Antecessores do kubernetes - Pesquisar depois por curiosidade)

## Referências

[KUBERNETES](http://kubernetes.io/pt/)
