## Intent
Memento는 구현을 몰라도 오브젝트의 이전 상태를 저장하고 복구시킬 수 있게 해준다.

## Problem
텍스트 에디터 앱을 만든다고 가정해보자. 만약 이전에 실행했던 행동을 취소하는 기능을 추가하려고 하는 경우, 앱은 모든 오브젝트의 스테이트들을 저장하고 있어야할 것이다. 매 순간 모든 오브젝트의 스테이트와 관련된 스냅샷을 생성할 경우 여러 문제가 발생하게 된다. 오브젝트들을 퍼블릭으로 만들면 안전하지 않고, 추후 개발이 진행되면서 액세스 접근권한 등 고려해질 사항이 너무 많아진다.

## Solution
Memento pattern delegates creting the snapshots to the actual owner of that state. Storing the copy of the object's tate in a special obejct called memento. Memento aren't accessible to any other object except the one that produced it.

## Structure
- Originator
  - Can produce snapshots of its own state. 
- Memento
  - Value object that acts as a snapshot of the originator's state. 
- Caretaker
  - knows not only 'when' and 'why' to capture the originator's state, but also when the state should be restored.

## How to Implement
1. Determine what class will play the role of the originator
2. Create the memento class
3. Make the memento class immutable.
4. Add a method for producing mementos to the originator class
5. Add a method for restoring the originator's state to its class
6. The caretaker should know when to request new mementos from the originator.
7. Link between caretakers and originators may be moved into the memento class.

## Pros and Cons
- Pros
  - You can produce snapshots of the object's state without violating its encapsulation
  - Simplify the originator's code by letting the caretaker maintain the history of the originator's state
- Cons
  - Lots of RAM usage
  - Caretaker should track the originator's lifecycle
  - Dynamic programming languages can't guarantee the state

## Code
```go run *.go```