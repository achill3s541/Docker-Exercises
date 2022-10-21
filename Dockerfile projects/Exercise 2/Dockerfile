FROM nginx:1.23.2
RUN apt-get update && apt-get install -y procps
COPY ./config/nginx.conf /etc/nginx/nginx.conf
#Creating html directory
RUN mkdir -p /var/www/html
COPY ./simple_website/index.html /var/www/html/index.html