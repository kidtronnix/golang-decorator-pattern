# golang-decorator-pattern

This is an example of the decorator pattern in golang. It is very useful when
a metho has many orthogonal concerns. What does that mean?
Say we wanna do a http request, pretty simple right. But what if we also need to perform some logging of the request
as well as add the requested action to an audit trail, etc. None of these concerns are related,
other than that we wanna do all of them whenever we call our method. This is where the decorator pattern really helps.

By implementing the decorator pattern we subscribe to the open-closed principal.
Our 'Do' method is open to extension but closed to modification. There's a lot of groovy benefits to obeying the open-closed pricniple.

For more info see: https://www.youtube.com/watch?v=xyDkyFjzFVc
