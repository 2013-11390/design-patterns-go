## Intent
Command는 리퀘스트를 stand-alone한 오브젝트로 만든다.

## Problem
새로운 텍스트 에디어 앱을 만든다고 가정하자. 여러 버튼이 있는 툴바를 만든다고 가정하였을 떄, 버튼 클래스를 만들어 사용한다. 이때, 여러 버튼들은 하나의 클래스에 속해 있지만 모두 역할은 다르다. 여러 서브 클래스들을 만들어 코드를 넣어서 해결할 수 있지만 이러한 경우 아래와 같은 문제점이 발생한다.
- 너무 많은 서브클래스들이 생성된다.
- 몇개의 동작(copy, paste)는 여러군데에서 쓰이는데 이러한 부분을 중복해서 작성하게 된다.

## Solution
GUI 와 비즈니스 로직을 분리한다. GUI레이어는 렌더링 관련해서만 동작시키고 비즈니스 로직은 분리한다. 커맨드 오브젝트는 GUI와 비즈니스 로직 오브젝트간의 연결역할을 한다. 이때, 커맨드들은 똑같은 인터페이스를 구성해야한다. 

## Structure
- Sender(Invoker)
  - 센더는 리퀘스트를 요청한다. 이 클래스는 커맨드 오브젝트의 레퍼런스를 갖고 있어야한다. 
- Command
  - 커맨드 인터페이스는 execute 메소드 하나를 정의한다
- Concrete Commands
  - 비즈니스 로직을 구성하는 것이 아니라 요청을 비즈니스 로직 오브젝트에 전달해주는 역할을 한다.
- Receiver
  - 비즈니스 로직을 갖고 있는다.

## How to Implement
1. Declare the command interface with a single execution method
2. Extract request into concrete command classes.
3. Sender 역할을 할 classes들을 정의한다. Sender들은 Commands와 오직 command interface를 통해 통신한다. 

## Pros and Cons
- Pros
  - Single Responsibility Principle
  - Open/Closed Principle
  - Can implement undo/redo
  - Can implement deferred execution of operations
  - Assemble a set of simple commands into a complex one.
- Cons
  - Code may be more complicated

## Code
go run *.go