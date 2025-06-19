import EndpointCard from "../components/EndpointCard";

export default function Docs() {
  return (
    <section className="p-6 max-w-5xl mx-auto">
      <h1 className="text-4xl font-bold mb-6">📚 Documentación de la API</h1>

      <EndpointCard
        method="GET"
        path="/api/v1/countries"
        title="Listado de países"
        description="Devuelve todos los países con filtros opcionales por región, subregión o nombre."
        example="/api/v1/countries?region=Asia&limit=10"
      />

      <EndpointCard
        method="GET"
        path="/api/v1/countries/{id}"
        title="Detalle de un país"
        description="Devuelve información completa de un país, incluyendo ciudades, monedas e idiomas."
        example="/api/v1/countries/1"
      />

      <EndpointCard
        method="GET"
        path="/api/v1/languages"
        title="Listado de idiomas"
        description="Devuelve todos los idiomas disponibles."
        example="/api/v1/languages"
      />

      <EndpointCard
        method="GET"
        path="/api/v1/cities?min_population=1000000"
        title="Ciudades por población"
        description="Filtra ciudades con más de cierto número de habitantes."
        example="/api/v1/cities?min_population=1000000"
      />

      <p className="mt-8 text-center text-gray-500">
        ⚙️ Más endpoints en desarrollo...
      </p>
    </section>
  );
}