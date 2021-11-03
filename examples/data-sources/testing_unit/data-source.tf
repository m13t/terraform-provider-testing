data "testing_unit" "static" {
  subject = "static tests"

  string_equal {
    statement = "test"

    input  = null
    expect = "test"
  }

  bool_equal {
    statement = "true should equal true"

    input  = true
    expect = true
  }

  bool_equal {
    statement = "testing another bool equal"

    input  = true
    expect = true
  }

  bool_not_equal {
    statement = "false should not equal true"

    input  = false
    expect = true
  }

  string_equal {
    statement = "X should equal X"

    input  = "X"
    expect = "X"
  }

  string_not_equal {
    statement = "X should not equal Y"

    input  = "X"
    expect = "Y"
  }

  integer_equal {
    statement = "10 should equal 10"

    input  = 10
    expect = 10
  }

  integer_not_equal {
    statement = "10 should not equal 20"

    input  = 10
    expect = 20
  }

  integer_less_than {
    statement = "10 should be less than 100"

    input  = 10
    expect = 100
  }

  integer_greater_than {
    statement = "100 should be greater than 10"

    input  = 100
    expect = 10
  }

  float_equal {
    statement = "3.142 should equal 3.142"

    input  = 3.142
    expect = 3.142
  }

  float_not_equal {
    statement = "3.142 should not equal 31.42"

    input  = 3.142
    expect = 3.142
  }

  float_less_than {
    statement = "3.142 should be less than 31.42"

    input  = 3.142
    expect = 31.42
  }

  float_greater_than {
    statement = "31.42 should be less than 3.142"

    input  = 31.42
    expect = 3.142
  }
}
