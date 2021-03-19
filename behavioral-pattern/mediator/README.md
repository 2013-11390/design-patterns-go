## Intent
Mediator를 이용하여 오브젝트간 디펜던시 혼잡을 해결해준다. 이 패턴은 오브젝트간 direct communication을 막고 mediator 오브젝트를 통해서만 협력할 수 있게 한다.

## Problem
고객 프로필 수정 페이지를 만든다고 가정해보자. 몇개의 element들은 다른 엘레먼트들과 interact하는 경우가 있다. 예를 들어, 체크박스를 클릭시 히든 텍스트가 보이는 경우 등이 있다. 이럴 때, 체크박스 코드 안에 히든 텍스트를 보이게 하는 코드를 넣을 시, 코드 재사용이 힘들어진다.

## Solution
Mediator 패턴은 컴퍼넌트와의 통신을 모두 mediator와 하게 한다. 위 예시의 경우, dialog 클래스가 각각의 엘레먼트들과의 통신을 중간에서 해준다.

## Structure
- Components
  - Various classes that contain some business logic
  - Each component has a reference to a mediator
- Mediator
  - Declares methods of communication with components. (일반적으로 single notification method)
- Concrete Mediator
  - Encapsulate relation between various components. 

## How to Implement
1. Identify a group of tighty coupled classes which would benefit from being more independent
2. Declare the mediator interface and describe communication protocol between mediators and various components.
3. Implement the concrete mediator

## Pros and Cons
- Pros
  - Single Responsbility Principle
  - Open/Closed Principle
  - Reduce coupling between components
  - Reuse individual components easily
- Cons
  - Mediator will become God Object

## Code
```go run *.go```