<<<<<<< HEAD

/*
language: go
go: 
 - 1.8.3
sudo: false

script:
 - go test -v ./...
*/
#!/bin/bash
node{
  stage('Dockerfile Envirionment'){
    docker image build -t rand .
//    go run ./main.go
    docker container run -p 4000:8080 --name randnum --rm rand 
  }
}
node {
  stage ('Unit Test') {
//    sh './go test -v ./..'
      sh 'go test'
  }
}
=======
node {
    echo 'Hello FreeWheel!'
}
>>>>>>> 922efbd991cce3ccadb290207f9b078aeb785282
