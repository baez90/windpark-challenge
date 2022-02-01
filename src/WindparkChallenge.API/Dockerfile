FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["src/WindparkChallenge.API/WindparkChallenge.API.csproj", "WindparkChallenge.API/"]
RUN dotnet restore "src/WindparkChallenge.API/WindparkChallenge.API.csproj"
COPY . .
WORKDIR "/src/WindparkChallenge.API"
RUN dotnet build "WindparkChallenge.API.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "WindparkChallenge.API.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "WindparkChallenge.API.dll"]
