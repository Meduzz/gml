## Attributes

This package contains a lot of the standard attributes used when creating them html stuff.

## Definition

"I noticed attributes are defined as `Attribute []string`.... wtf?"

Yes. Inspired by slog, where everything is `...string`. But when you start creating helpers for attributes, they return `[]string`. So then a choice had to be made. Wrap all helpers in another helper that turns them back to `[]string` or believe in the attribute helpers and make them first class citizen and create a helper for `...string` style attributes.