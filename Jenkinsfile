pipeline {
  agent any
  stages {
    stage('Checkout') {
      steps {
        git(url: 'https://github.com/dhamodharanv-workhall/jenkins-lab', branch: 'main')
      }
    }

    stage('List Files') {
      steps {
        sh 'ls -la'
      }
    }

  }
}