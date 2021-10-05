docker build -t nitrodenov/otus:v7.2 . docker push nitrodenov/otus:v7.2

helm install pg bitnami/postgresql -f postgresql_values.yaml

helm install otus ./otus-app 

kubectl apply -f app-config.yaml

Для достижения идемпотентности в order используется версионирование списка заказов.
При создании заказа в запросе передается заголовок If-Match: <версия списка>. 
В случае если в запросе передается последняя версия, то заказ успешно создается с новой версией (+1 к версии переданной в заголовке If-Match). 
В противном случае возвращается ответ со статусом 409 Conflict и заголовком ETag: <последняя версия заказа>.
