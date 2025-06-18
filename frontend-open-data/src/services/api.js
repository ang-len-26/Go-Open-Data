const BASE_URL = "http://localhost:8080/api/v1";

export const fetchCountries = async (params = {}) => {
  const query = new URLSearchParams(params).toString();
  const res = await fetch(`${BASE_URL}/countries?${query}`);
  if (!res.ok) throw new Error("Error al obtener datos");
  return res.json();
};
