## Intent
Chain of Responsiblity는 리퀘스트가 여러 핸들러들을 통해 처리되는 것을 의미한다. 각각의 핸들러들은 리퀘스트를 처리하고 다음 핸들러로 넘길지 판단한다.

## Problem
주문 시스템에서 유저들의 인증이 처리된 후에 주문 시스템에 요청이 들어가야한다. 만약 인증처리가 되지 않으면 이후에 처리는 진행이 될 필요가 없다. 이러한 과정이 여러개가 될 수 있다.
- 주문 시스템에 Raw data를 쓰면 보안상 취약해 validation step을 추가한다.
- 시스템이 brute force attack에 취약하여 repeated failure에 대한 체크 시스템을 추가한다.
- Repeated request에 대한 리스폰스를 위한 캐쉬를 추가한다.

## Solution
Chain of Responsibility는 handler가 각각의 역할을 실행한다. 각 handler에 적합한 데이터가 들어오는 경우 다음 handler에게 리퀘스트를 날려 진행하게 된다. 이 경우, handler들은 모두 공통적인 인터페이스를 갖고 있어야 한다.
- Handler
  - Interface 정의
- Base Handler
  - Optional boilerplate class for all common handler classes
- Concrete Handler
  - 리퀘스트를 처리하기 위한 실제 코드 작성

## How to Implement
1. Declare the handler interface
2. Create an abstract base handler class. 이 클래스는 다음 핸들러를 저장하기 위한 필드가 있어야 한다.
3. Create concrete handler subclasses. 모든 핸들러들은 두가지 사항을 고려해야한다.
  - 리퀘스트를 처리할 것인지
  - 리퀘스트를 패스시킬것인지

## Pros and Cons
- Pros
  - 리퀘스트 핸들링 순서를 정할 수 있다
  - Signle Responsibility Principle.
  - Open/Closed Principle. 새로운 핸들러를 앱에 쉽게 추가할 수 있다.
- Cons  
  - 몇 리퀘스트들을 unhandled된 채로 처리될 수 있다.