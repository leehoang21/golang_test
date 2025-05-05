docker run -d \
  --name mongo \
  -e MONGO_INITDB_ROOT_USERNAME=docker \
  -e MONGO_INITDB_ROOT_PASSWORD=docker \
  -p 27017:27017 \
  mongo:6
