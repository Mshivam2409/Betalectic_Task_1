cd auth/
docker build -t shivammalhotra/betalectic_task_1_auth .
docker push shivammalhotra/betalectic_task_1_auth
cd ../secure
docker build -t shivammalhotra/betalectic_task_1_secure . 
docker push shivammalhotra/betalectic_task_1_secure
cd ../test
docker build -t shivammalhotra/betalectic_task_1_test . 
docker push shivammalhotra/betalectic_task_1_test
cd ../db
docker build -t shivammalhotra/betalectic_task_1_mongo . 
docker push shivammalhotra/betalectic_task_1_mongo