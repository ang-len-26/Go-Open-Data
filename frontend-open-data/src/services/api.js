const BASE_URL = "http://localhost:8080/api/v1";

export const fetchCountries = async (params = {}) => {
  const query = new URLSearchParams(params).toString();
  const res = await fetch(`${BASE_URL}/countries?${query}`);
  if (!res.ok) throw new Error("Error al obtener países");
  return res.json();
};

export const fetchCities = async (params = {}) => {
  const query = new URLSearchParams(params).toString();
  const res = await fetch(`${BASE_URL}/cities?${query}`);
  if (!res.ok) throw new Error("Error al obtener ciudades");
  return res.json();
};

export const fetchLanguages = async () => {
  const res = await fetch(`${BASE_URL}/languages`);
  if (!res.ok) throw new Error("Error al obtener idiomas");
  return res.json();
};

export const fetchCurrencies = async () => {
  const res = await fetch(`${BASE_URL}/currencies`);
  if (!res.ok) throw new Error("Error al obtener monedas");
  return res.json();
};

