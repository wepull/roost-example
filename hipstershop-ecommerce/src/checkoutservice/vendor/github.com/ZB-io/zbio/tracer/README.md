# Tracing package

## Expected functionality

- Create NewTracer
  - with unique Name
  - with StartOption
    - pass kv attriutes
    - linkTo(span)
- Get span := SpanFromContext(ctx)     // parse ctx.Values().(whichType)
  - parse which type of data is in context and based on that return span handler

## References
