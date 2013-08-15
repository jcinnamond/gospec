/*
Package gospec provides a spec like syntax for writing Go tests.

Use Verify() and Equals() to build simple equality tests.

    func TestUser_Fullname(t testing.T) {
        user := User.new("John", "Cinnamond")
 	    if err := gospec.Verify(user.Fullname()).Equals("John Cinnamond"); err != nil {
	  	  t.Error(err)
	    }
    }

You can also add DoesNot() to the middle of the call to check for inequality.

    func TestCount(t testing.T) {
        c = counter.New()
        c.Inc()
 	    if err := gospec.Verify(c.NextVal()).DoesNot().Equal(0); err != nil {
	  	  t.Error(err)
	    }
    }

*/
package gospec
