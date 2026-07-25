# web

TypeScript source for the viewer SPA. Built by Vite into `../internal/viewer/dist`
and compiled into the Go binary under the `spa` build tag.

There is no build-time coupling in the other direction. Nothing here knows about
Go, and `go build ./...` without the `spa` tag does not run this build.

## Commands

Run from this directory, or use `make web` from the repo root.

| Command | Effect |
| --- | --- |
| `npm run build` | Type check, then emit into `../internal/viewer/dist` |
| `npm run dev` | Vite dev server, proxying `/api` to a running `joist viewer` |
| `npm run check` | Type check only |

## Intended views

Recorded here so the scope stays fixed. None of it is built yet.

- Design structure matrix with partitioning, as the primary view.
- Ranked violation list, ordered by the cost of severing the edge.
- Trend line of violation count across commits.
- Disagreement quadrant: high co-change against no static edge.

Explicitly not a force-directed node-link diagram. At this graph size that is a
hairball and communicates nothing.

## Data

The SPA is read only. It reads one artifact from `GET /api/artifact` and never
writes. The artifact format is the `schema` Go module; keep any TypeScript types
here in sync with it by hand until there is a reason to generate them.
