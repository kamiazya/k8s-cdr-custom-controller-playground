# k8s-cdr-custom-controller-playground

Playground to create k8s CDR and Custom Controller.

## Commands

### Apply CDR

```bash
$ kubectl apply -f crd.yml
customresourcedefinition.apiextensions.k8s.io/hoges.sample.kamiazya.tech created
```

### Code Generation

#### Before run

```bash
go get k8s.io/client-go/...
go get k8s.io/apimachinery/...
go get github.com/golang/glog
go get k8s.io/code-generator/...
```

#### Generate

```bash
# k8s.io/code-generator/generate-groups.sh
$ ../../../k8s.io/code-generator/generate-groups.sh \
  all \
  github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/client \
  github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/apis \
  hoge:v1
Generating deepcopy funcs
Generating clientset for hoge:v1 at github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/client/clientset
Generating listers for hoge:v1 at github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/client/listers
Generating informers for hoge:v1 at github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/client/informers
```
