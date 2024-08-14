#/bin/bash

# loop infinite

while true; do
  # curl "http://ip/api/users-service/users"
  # curl "http://ip/api/books-service/books"
  curl "http://books-api.local/api/users-service/users"
  curl "http://books-api.local/api/books-service/books"
  #   sleep 1
done

