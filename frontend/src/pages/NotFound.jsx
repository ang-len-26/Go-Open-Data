export default function NotFound() {
  return (
    <div className="h-screen flex flex-col justify-center items-center">
      <h1 className="text-5xl font-bold">404</h1>
      <p className="text-gray-600 mt-2">Página no encontrada</p>
      <Link to="/" className="mt-4 text-blue-600 hover:underline">
        Volver al inicio
      </Link>
    </div>
  );
}