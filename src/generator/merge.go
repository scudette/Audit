package generator

// Merge rules from the other into this model
func (self *Rules) Merge(other *Rules) error {
	// First build an index for fast lookup
	index := make(map[string]*Check)
	for _, o := range other.Checks {
		index[o.Id] = o
	}

	for _, c := range self.Checks {
		old_check, pres := index[c.Id]
		if !pres {
			continue
		}

		c.Merge(old_check)
	}

	return nil
}

func (self *Check) Merge(other *Check) error {
	self.Verified = other.Verified

	index := make(map[string]*Test)
	for _, t := range other.Rules {
		index[t.OriginalTest] = t
	}

	for _, t := range self.Rules {
		old_test, pres := index[t.OriginalTest]
		if pres {
			if old_test.Type != "" {
				t.Type = old_test.Type
			}

			if old_test.Name != "" {
				t.Name = old_test.Name
			}

			if old_test.ColumnExpression != "" {
				t.ColumnExpression = old_test.ColumnExpression
			}

			if old_test.WhereExpression != "" {
				t.WhereExpression = old_test.WhereExpression
			}

			if old_test.Env != nil {
				t.Env = old_test.Env
			}
		}
	}

	return nil
}
