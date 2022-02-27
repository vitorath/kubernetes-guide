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

## Acessando API's interna no Kubernetes

A fim de curiosidade para visualizar todas as api's do kubernetes

```sh
kubectl proxy --port=8080 # No caso foi exposto na porta 8080, mas poderia ser outra
```

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

**ReplicaSet** gerencia os **pods** mantendo em execução o número especificado de replicas. Contudo, quando a versão do **pod** é atualizada o **ReplicaSet não recria os pods automaticamenteo**, ou seja, somente novos pods estarão com a versão mais atualizada da aplicação.

O arquivo (k8s/kube-replicaset)[./k8s/kube-replicaset/] é um exemplo de ReplicaSet e para executa-lo utilize o **apply**, similar ao pod.

```sh
# Sendo "k8s/kube-replicaset.yaml" o caminho do arquivo de configuração
kubectl apply -f k8s/kube-replicaset.yaml
```

Para visualizar o(s) ReplicaSets criado execute:

```sh
kubectl get replicasets
```

kubectl delete replicaset goserver-replicaset

## Deployment

Um **Deployment** fornece atualizações declarativas para os **Pods** e os **ReplicaSets**, isto é, ao atualizar um arquivo de configuração de **Deployment** o kubertentes cria um novo **ReplicaSet** e aos poucos vai derrubando os os **Pods** do **ReplicaSet** antigo e criando novas **Pods** com o novo **ReplicaSet**, gerando assim **zero downtime** (a aplicação não fica fora do ar). Abaixo o comando que aplicado as configurações do arquivo (k8s/kube-deployment.yaml)[k8s/kube-deployment.yaml].

```sh
# Sendo "k8s/kube-deployment.yaml" o caminho do arquivo
kubectl apply -f k8s/kube-deployment.yaml
```

## Rollout

**Rollout** consiste na mudança de versões de objetos já criados. Por exemplo para voltar à versão anterior do **deployment**

```sh
# Sendo "deployment" o objeto e "goserver-deployment" o nome
kubectl rollout undo deployment goserver-deployment
```

Agora para voltar à uma versão específica

```sh
# Sendo "deployment" o objeto e "goserver-deployment" o nome e "2" o número da revisãO
kubectl rollout undo deployment goserver-deployment --to-revision=2
```

Para verificar o histórico de rolluot e obter o número da revisão

```sh
# Sendo "deployment" o objeto e "goserver-deployment"
kubectl rollout history deployment goserver-deployment
```

## Service

Service tem como proprósito de expor sua aplicação automaticamente e já aplicando a um processo de service discovery.

Para listar todos os services

```sh
kubectl get service
kubectl get svc # Abreviação
```

### ClusterIP

Consiste em mapear um ip interno para acessar o service.

Configurando uma aplicação que está sendo executada na porta 8000 (**targetPort**) e exposta no serivce pela porta 80 (**port**), além de forçar o acesso externo na porta 9000. Em resumo, o usuário acessa a porta 9000 que é redirecionada para a porta 80 que posteriomente acessa a porta 8000.

```sh
kubectl apply -f k8s/kube-service-cluster-ip.yaml
kubectl port-forward svc/goserver-service 9000:80
```

### NodePorts

NodePort é a forma mais arcaica para acessar um cluster kubernetes de forma externa. Neste serviço o kuberntes libera a mesma porta para todos os **nodes** do cluster. Além disso, a porta exposta deve estar entre 30000 e 32767.

Normalmente utilizado para quando for configurar o próprio **loadBalancer** ou ainda expor uma porta temporariamente.

Neste exemplo, o serviço está exposto na porta 30001 (**nodePort**) que acessa a porta 80(**Service**) que redireciona para a porta 8000 (**targetPort**).

```sh
kubectl apply -f k8s/kube-service-node-port.yaml
```

### LoadBalancer

LoadBalancer tem como propósito disponibilizar o acesso externo para á suas aplicações no kubernetes, gerando uma IP e uma porta externa automaticamente.

Normalmente este serviço é utilizado é um cluster gerenciado, isto é, em um provedor de nuvem.
Gera um Ip para ser acessado externamente almem de gerar um porta externa. Normalmente utilizado quando está sendo executado em um cluster gerenciado, isto é, um provedor de nuvem. Contudo com o **kind** a porta externa fica como pendente, pois normalmente sua máquina não está configurado todos os drivers de nuvem necessário para se auto configurar.

```sh
kubectl apply -f k8s/kube-service-load-balancer.yaml
```

## Variáveis de ambiente

Para adicionar variáveis de ambiente em sua aplicação, de uma forma simplificada, adicione diretamente a variável em um **pod**, conforme desmontrado no abaixo.

```sh
kubectl apply -f k8s/kube-deployment-env.yaml
```

Contudo existe uma forma mais interessante de fazer a mesma configuração e está consiste em utilizar um objeto de configuração denominado **ConfigMap**.

O **ConfigMap** é um objeto de configuração que pode ser utilizado em um pod. Uma das formas que este pode ser utilizado é para configurar variáveis de ambiente.

```sh
kubectl apply -f k8s/kube-configmap-env.yaml # Aplicando configmap
kubectl apply -f k8s/kube-deployment-env-configmap.yaml # Executando o deployment.
```

**Ps.:** as mudanças de fato serão efetivadas somente quando atualizar o **deployment**, caso modifique e atualize **somente o ConfigMap** as alterações **não serão aplicadas**.

## Criando arquivos com ConfigMap

Atém de configurar variáveis de ambiente, com o **ConfigMap** também é possível utiliza-lo para criar arquivos utilizados por sua aplicação. Neste cenário é necessário configurar um **volume** no qual o **kubernetes** irá aplicar as configurações do **ConfigMap** e gerar um arquivo.

é possível utilizar uma arquivo configMap como um arquivo físico
pega um configmap e transforma em voluma

```sh
kubectl apply -f k8s/kube-configmap-family-file.yaml # ConfigMap que simula um arquivo
kubectl apply -f k8s/kube-deployment-file-configmap.yaml # Aplicar configurações à aplicação
```

###  Secrets

As configurações de **secrets** devem ser encriptografadas em Base64. Posteriormente as secretes são configuradas como variáveis de ambiente do **container**.

## Probes

**Health check** é uma forma para identificar se uma aplicação está funcionando corretamente ou não. O **kubernetes** contempla com uma implementação nativa denomidade **Liveness Probe** que verifica de tempos em tempos a aplicação, desde que seja configurado, caso ocorra alguma problema o **kubernetes** automaticamente reinicia o **pod**.

Por outro lado, além de validar se a aplicação está funcionando é comum desviar o fluxo de dados daquele container que está com problema o mais rápido possível, pois dependendo a aplicação pode ocasionar um grande prejuizo para empresa. Para este propósito existe o **Readiness Probe** que tem como objetivo verificar se a aplicação está funcionado e consequentemente se houver algum problema irá derrubar a aplicação.

**Obs.:** Ao utilizar o **Liveness Probe**, considere configura o timeout da requisição devidamente, pois caso esteja fazendo algum **check integrado**(verificando banco por exemplo) considere o tempo que a aplicação demora para fazer todos estes processos. Além disso, é muito importante que o **Liveness Probe** e o **Readiness Probe** esteja perfeitamente sincronizados para não ocasionar um loop na criação dos containers e com isso a aplicação nunca estara funcionando.

### Heath Check 
Mecanismo que verifica de tempos em tempos a aplicação para verificar se está funcionando corretamente, no caso do kubernetes normalemnte reinicia o container.

Redirecionar o trafego somente quando a aplicação for inicializada

### Liveness probe
Verifica de tempos em tempos para verificar a saude da aplicação

Normalente o timeout configurado é um segundo
Contudo caso queira fazer um teste mais integrado, vale a pena fazer um healthz somente para isto e aumente o timeout

# Informações que podem ser utilizadas

Deployment > ReplicaSet > Pod

Borg > Omega > Kubernetes (Antecessores do kubernetes - Pesquisar depois por curiosidade)
kubectl describe pod <nome-pod>
kubectl describe deployment <nome-deployment>

kubectl delete svc <nome-service>

kubectl exec -t <nome-pod> -- <comando>
kubectl logs <nome-pod>


➜ kubectl apply -f k8s/kube-deployment-check.yaml && watch -n1 kubectl get po


## Referências

[KUBERNETES](http://kubernetes.io/pt/)
