export default function EndpointCard({ title, method, path, description, example }) {
  return (
    <div className="border p-5 rounded-xl shadow bg-white mb-6">
      <div className="text-sm text-gray-600 font-mono">
        <span className="text-green-600 font-bold mr-2">{method}</span>
        {path}
      </div>
      <h2 className="text-xl font-semibold mt-2">{title}</h2>
      <p className="mt-1 text-gray-700">{description}</p>
      {example && (
        <pre className="mt-3 bg-gray-100 p-3 rounded text-sm text-gray-800">
          {example}
        </pre>
      )}
    </div>
  );
}