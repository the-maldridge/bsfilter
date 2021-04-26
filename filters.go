package bsfilter

// NewExpressionSet provisions a new expressionset.
func NewExpressionSet() *ExpressionSet {
	return &ExpressionSet{
		expressions: make(map[string]*Expression),
	}
}

// Filter runs all expressions in the set over the provided ValueSet
// and returns a []string of expressions that matched.
func (es *ExpressionSet) Filter(v ValueSet) []string {
	es.RLock()
	out := []string{}
	for f, e := range es.expressions {
		if e.Evaluate(v) {
			out = append(out, f)
		}
	}
	es.RUnlock()
	return out
}

// Add inserts a new expression into the set.  The name must be unique
// and inserting the same name twice will overwrite entries.
func (es *ExpressionSet) Add(name string, e *Expression) {
	es.Lock()
	es.expressions[name] = e
	es.Unlock()
}

// Del removes an expression from the set.
func (es *ExpressionSet) Del(name string) {
	es.Lock()
	delete(es.expressions, name)
	es.Unlock()
}
