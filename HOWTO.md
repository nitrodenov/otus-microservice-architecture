docker build -t nitrodenov/otus:v5.1.0 .
docker push nitrodenov/otus:v5.1.0

docker build -t nitrodenov/otus:v5.2.0 .
docker push nitrodenov/otus:v5.2.0

helm install auth-app ./auth-chart
helm install user-info-app ./user-chart

kubectl apply -f app-config.yaml 

helm install pg bitnami/postgresql -f postgresql_values.yaml