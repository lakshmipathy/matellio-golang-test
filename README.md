# matellio-golang-test
Rest micro services 

git clone https://github.com/go-lakshmi/matellio-golang-test.git
install docker 
cd matellio-golang-test
docker build -t mattest:latest .
docker image ls
docker run -p 0.0.0.0:8080:8080 {docker-image-id}

# Rest Api details
# get all ports 
# HTTP METHOD GET 
http://localhost:8080/api/ports

# get by port id 
# HTTP METHOD GET 
http://localhost:8080/api/port/{id}
http://localhost:8080/api/port/AEJEA

# Create or Update Port 
# HTTP METHOD POST
http://localhost:8080/api/port/{id}
http://localhost:8080/api/port/AEJEA
# request body
{
        "userName": "",
        "city": "Jebel Ali",
        "country": "United Arab Emirates",
        "alias": [],
        "regions": [],
        "coordinates": [
            "",
            ""
        ],
        "province": "Dubai",
        "timezone": "Asia/Dubai",
        "unlocs": [
            "AEJEA"
        ],
        "code": "52051"
    }

# Update Port  
# HTTP METHOD PUT
http://localhost:8080/api/port/{id}
http://localhost:8080/api/port/AEJEA
# request body
{
        "userName": "",
        "city": "Jebel Ali",
        "country": "United Arab Emirates",
        "alias": [],
        "regions": [],
        "coordinates": [
            "",
            ""
        ],
        "province": "Dubai",
        "timezone": "Asia/Dubai",
        "unlocs": [
            "AEJEA"
        ],
        "code": "52051"
    }
# delete Port  
# HTTP METHOD DELETE
http://localhost:8080/api/port/{id}
http://localhost:8080/api/port/AEJEA     


# Notes
1) Implemented Github Actions CI-CD pattern to build and execute LINT Test, this we can extend to BUILD, TEST and DEPLOY a  complete CI-CD.
2) Used rest api for testing convinience. which easily can access the api endpoints easily
3) As part of Api testing not handled the validation, negative scanarios, load testing, Performance testing and Limit Test
4) Regarding loading big json file. need to push records to database, instead of storing in in-memory
 
