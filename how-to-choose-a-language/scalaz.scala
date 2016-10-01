import scalaz._
import scalaz.std.list._
import scalaz.syntax.monad._
import scalaz.syntax.monoid._
import scalaz.syntax.traverse.{ToFunctorOps => _, _}

class Foo[F[+_] : Monad, A, B](val execute: Foo.Request[A] => F[B], val joins: Foo.Request[A] => B => List[Foo.Request[A]])(implicit J: Foo.Join[A, B]) {

def bar: Foo[({type l[+a]=WriterT[F, Log[A, B], a]})#l, A, B] = {
    type TraceW[FF[+_], +AA] = WriterT[FF, Log[A, B], AA]
    def execute(request: Request[A]): WriterT[F, Log[A, B], B] =
      self.execute(request).liftM[TraceW] :++>> (repr => List(request -> request.response(repr, self.joins(request)(repr))))
  ----- REDACTED -------
  
  
