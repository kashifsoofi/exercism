/// What should the type of _function be?
pub fn map<TIn, TOut, F>(input: Vec<TIn>, function: F) -> Vec<TOut> where F: FnMut(TIn) -> TOut {
    input.into_iter().map(function).collect()
}
