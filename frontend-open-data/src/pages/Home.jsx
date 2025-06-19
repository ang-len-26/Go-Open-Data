import { Link } from "react-router-dom";

export default function Home() {
  return (
    <section className="p-8 max-w-4xl mx-auto text-center">
      <h1 className="text-5xl font-bold mb-6">🌐 Open Data API</h1>
      <p className="text-lg text-gray-700 mb-8">
        API pública que expone datos estructurados sobre países, regiones,
        idiomas, monedas y más. Ideal para visualizaciones, dashboards y
        proyectos educativos.
      </p>
      <Link
        to="/docs"
        className="inline-block bg-blue-600 text-white px-6 py-3 rounded-lg shadow hover:bg-blue-700"
      >
        📘 Ver documentación
      </Link>
    </section>
  );
}
