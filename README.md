Minimales Beispiel für eine Go-Applikation in OpenShift:
(läuft auf jedem OpenShift Cluster)

# Features
Minimale Go-Applikation mit `html/template` und Layout in Bootstrap.   

OpenShift erstellt automatisch eine BuildConfig und DeploymentConfig (etc.), baut ein Docker-Image in der internen Registry und lässt es laufen.  


# OpenShift

Erstellen in OC:

    oc new-project go
    oc project go  
    oc new-app centos/go-toolset-7-centos7~https://github.com/rschumm/hallogo.git


Bemerkung: die language-Detection von `new-app` für OpenShift 3.11 scheint nicht korrekt zu laufen, darum benutze ich "von Hand" dieses [s2i Builder Image](https://github.com/sclorg/golang-container).  
Dieses Image erstellt keinen Service, darum muss der noch nach-erstellt werden: 

    oc expose dc/demoapp-go --port=8080

Erstelle jetzt eine Route zum generierten Service - und die App sollte erreichbar sein. Alternativ auf der Kommantozeile:

    oc expose svc/demoapp-go

alles wieder loswerden:

    oc delete all --selector app=hallogo   

# Local run

die Applikation local laufenlassen: 

    go run main.go

und [http://localhost:8080/index.html](http://localhost:8080/index.html) aufrufen. 

