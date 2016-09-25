package euler
import "testing";

func TestProblem(t *testing.T) {
   if (findSum() != 4613732) {
     t.Error("Incorrect result, expected 4613732");
   }
}
