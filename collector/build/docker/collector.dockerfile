FROM gcr.io/distroless/static:nonroot

USER nonroot:nonroot

COPY --chown=nonroot:nonroot wp /app/

ENTRYPOINT ["/app/wp"]
CMD ["collect"]