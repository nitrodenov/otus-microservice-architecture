1) команда установки БД из helm, вместе с файлом values.yaml.
helm install pg bitnami/postgresql -f postgresql_values.yaml

2) команда применения первоначальных миграций
kubectl apply -f initdb.yaml
   
3) команда kubectl apply -f, которая запускает в правильном порядке манифесты кубернетеса
kubectl apply -f app-config.yaml
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml

4) Postman коллекция, в которой будут представлены примеры запросов к сервису на создание, получение, изменение и удаление пользователя.
   https://www.getpostman.com/collections/7722c81a11ce3a1298ef