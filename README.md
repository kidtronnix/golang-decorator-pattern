# golang-decorator-pattern

This is an example of the decorator pattern in golang. It is very useful when
A method or action has a bunch of orthogonal concerns. What does that mean?
So say we wanna do a http request AND perform some logging of this request
AND add this action to an audit trail. None of these concerns are related,
other than that we wanna do all of them whenever we call Do(r)

By implementing the decorator pattern we subscribe to the open-closed principal.
In that our 'Do' method is open to extension but closed to modification.
For more info see: https://www.youtube.com/watch?v=xyDkyFjzFVc
