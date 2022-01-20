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

# Kind - uma ferramenta para praticar com Kubernetes

[Kind](https://kind.sigs.k8s.io/docs/user/quick-start) é uma ferramenta que executa localmente clusters de Kubernetes usando "nodes" containers Docker. Contudo, para ter acesso aos clusters que são gerandos pelo **kind** é mais conveniente utilizar o [kubectl](https://kubernetes.io/docs/tasks/tools/) que é client do kuberntes que se comunica com o servidor do kubernetes.

## Criando clusters com o Kind

Para criar um cluster com um único **node** utilizando o kind execute o comando:

```sh
kind create cluster
```

Contudo, existe outra forma de criar clusters no kind por meio de um arquivos de configuração, neste caso [k8s/kind-config.yaml](./k8s/kind-config.yaml)

```sh
kind create cluster --config=k8s/kind-config.yaml
```

Após criar o cluster é necessário conectar nele, ou seja, indicar qual o contexto da conexão. No Linux, as credenciais de conexão do **kubectl** podem ser encontrada neste caminho **~/.kube/config**, lembrando que esta pasta somente irá existir a partir do momento que houver um conexão de um cluster configurada. Agora para efetuar o acesso ao cluster digite:

```sh
kubectl cluster-info --context kind-kind
```

Para verificamos se estamos conectados o comando abaixo deve retornar um ou mais nodes

```sh
kubectl get nodes
```

Contudo caso, sejá necessário deletar um cluster do kind devemos saber qual o nome do cluster (por padrão, caso não tenha especificado o nome é kind) e posteriormente executar o comando de delete

```sh
kind get clusters # Lista todos os clusters criados pelo kind
kind delete clusters kind # Deleta o cluster com o nome "kind"
```

# Informações que podem ser utilizadas

Borg > Omega > Kubernetes (Antecessores do kubernetes - Pesquisar depois por curiosidade)

## Referências

[KUBERNETES](http://kubernetes.io/pt/)
