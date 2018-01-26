### 윈도우즈 10 에 파이키덤 설치 실패로 넘어감

### 데이터 버저닝의 필요성.

협력 : 내가 만든 것을 타인의 리뷰를 받으려면 저장이 되어 있어야 한다.

창의성 : 내 기억력도 못 믿는다. 저장된 우리 모델과 테크닉으로부터 계속적 계선이 필요하다.

법규 : 자료에 관한 법규가있는데.. "무슨 소리지?" 우리가 무엇을 처리하고 결과를 생산하는지 추적할 필요가 있다. 


pachyderm 은 데이터 버저닝 도구.

- 고 클라이언트가 있다.
- 어떤 타입과 어떤 자료의 형식이건 버전이 가능하다.
- 버전된 자료를 보관하는 유연한 객체 저장소.
- 운전중인 버전된 ML 워크 플로우를 위한 데이터 파이프라이닝 시스템과의 통합.

### 파이키덤 자르곤

깃이랑 비슷하다.

- 저장소: 버전된 자료의 컬렉션
- 커밋 : 자료 저장소에 자료 커밋을 만듦으로써 버전이 된다.
- 브랜치 : lightweight 는 어떤 커밋들이나 커밋들의 묶음을 가리킨다.
- 파일들 : 자료는 파이키덤에서 파일레벨로 버전된다, 
    그리고 파이키덤은 자동적으로 전략을 가동하는데, 중복제거 등, 버전된 자료의 공간을 효율적으로 유지한다.
   
차이점이 있다. 페타바이트 자료를 합칠 때, 사람은 하지 못할 것이다.

### 파이키덤 배포/설치

왜이렇게 복잡하냐. 파이키덤은 쿠버니츠에서 돌아간다.
쿠버니츠는 컨테이너화된 어플리케이션을 자동 배포, 스케일링, 관리를 하게 해주는 시스템이다.

minikube 는 싱글 큐버니트.

1.  kubectl [설치](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
   ```
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    sudo mv ./kubectl /usr/local/bin/kubectl
   ``` 
2. minikube [설치](https://github.com/kubernetes/minikube/releases)
   윈도우즈에서 설치하는 [방법](https://gist.github.com/williballenthin/8bbb7710ca8faa93d6eb57a44eda785b)
   `curl -Lo minikube https://storage.googleapis.com/minikube/releases/v0.24.1/minikube-linux-amd64 && chmod +x minikube && sudo mv minikube /usr/local/bin/`
