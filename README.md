# DCM Docker Compose smoke test (web + redis)

A minimal two-service Compose project for exercising DCM's multi-container
deployments.

- `web` — a tiny Go HTTP server (built from `Dockerfile`) that PINGs the `redis`
  service over the Compose network and responds on `:8080`.
- `redis` — the stock `redis:alpine` image.

## Use it with DCM

1. Push these four files to a new HTTPS Git repository (e.g.
   `https://github.com/<you>/dcm-compose-demo.git`), default branch `main`.
2. In DCM: **Create Project** with that Git URL and branch `main`. The internal
   port is ignored for Compose projects.
3. **Deploy**. DCM detects `compose.yaml`, validates it, synthesizes its own
   hardened compose file, and brings up both services on a per-deployment
   private network.
4. Open the **Containers** tab: `web` and `redis` should both be `running`.
   Follow the `web` host port link — it should say
   `Hello from DCM — redis PONG ok`.
5. **Logs** per service, **Restart** an individual service, then **Redeploy** —
   the previous deployment stays up until the new one is healthy, then is
   marked `deployed`.

hola we 