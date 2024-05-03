~Veiga

My idea:

- `pkg/`: Libraries. Could have been `lib/`, but apparently `pkg/` is more goy (?).
- `pkg/server-common/`: Server logic that is common for all server implementations: The cloud one using lambdas and the on-premise one for LAN. Honestly it's kind of weird for it to be buried here since it is probably the main thing...
- `client/`: Game logic that runs on the client. Should display graphics, collect inputs, etc.
- `server-lambda/`: Server implementation using AWS lambdas. Requires `pkg/server-common/`.
