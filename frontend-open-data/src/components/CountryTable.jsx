import { useState, useEffect } from "react";
import { fetchCountries } from "../services/api";

export default function CountryTable() {
  const [countries, setCountries] = useState([]);
  const [filters, setFilters] = useState({ region: "", subregion: "", name: "" });
  const [pagination, setPagination] = useState({ limit: 10, offset: 0 });
  const [total, setTotal] = useState(0);

  const getData = async () => {
    try {
      const params = { ...filters, ...pagination };
      const data = await fetchCountries(params);
      setCountries(data.data);
      setTotal(data.total);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    getData();
  }, [filters, pagination]);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">Lista de Países</h1>

      <div className="mb-4 grid grid-cols-1 md:grid-cols-3 gap-2">
        <input
          className="p-2 border"
          placeholder="Región"
          value={filters.region}
          onChange={(e) => setFilters({ ...filters, region: e.target.value })}
        />
        <input
          className="p-2 border"
          placeholder="Subregión"
          value={filters.subregion}
          onChange={(e) => setFilters({ ...filters, subregion: e.target.value })}
        />
        <input
          className="p-2 border"
          placeholder="Nombre"
          value={filters.name}
          onChange={(e) => setFilters({ ...filters, name: e.target.value })}
        />
      </div>

      <table className="w-full border">
        <thead>
          <tr className="bg-gray-100">
            <th className="p-2 border">#</th>
            <th className="p-2 border">Nombre</th>
            <th className="p-2 border">Capital</th>
            <th className="p-2 border">Región</th>
			<th className="p-2 border">Subregión</th>
            <th className="p-2 border">Población</th>
            <th className="p-2 border">Área</th>
          </tr>
        </thead>
        <tbody>
          {countries.map((c) => (
            <tr key={c.id}>
              <td className="p-2 border">{c.id}</td>
              <td className="p-2 border">{c.name}</td>
              <td className="p-2 border">{c.capital}</td>
              <td className="p-2 border">{c.region}</td>
			  <td className="p-2 border">{c.subregion}</td>
              <td className="p-2 border">{c.population.toLocaleString()}</td>
              <td className="p-2 border">{c.area.toLocaleString()} km²</td>
            </tr>
          ))}
        </tbody>
      </table>

      <div className="flex justify-between mt-4">
        <button
          onClick={() =>
            setPagination((prev) => ({ ...prev, offset: Math.max(prev.offset - prev.limit, 0) }))
          }
          disabled={pagination.offset === 0}
          className="px-4 py-2 bg-blue-500 text-white disabled:opacity-50"
        >
          Anterior
        </button>
        <button
          onClick={() =>
            setPagination((prev) => ({ ...prev, offset: prev.offset + prev.limit }))
          }
          disabled={pagination.offset + pagination.limit >= total}
          className="px-4 py-2 bg-blue-500 text-white disabled:opacity-50"
        >
          Siguiente
        </button>
      </div>
    </div>
  );
}
