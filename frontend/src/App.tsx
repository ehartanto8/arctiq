import { useState } from 'react';

function App() {
    const [prompt, setPrompt] = useState('');
    const [tasks, setTasks] = useState<any[]>([]);

    return (
        <div className="p-6 max-w-xl mx-auto space-y-4">
            <h1 className="text-3xl font-bold">Arctiq</h1>

            <textarea
                className="w-full p-3 border rounded"
                rows={4}
                placeholder="What do you want to build?"
                value={prompt}
                onChange={(e) => setPrompt(e.target.value)} />

            <button className="px-4 py-2 bg-blue-600 text-white rounded">
                Generate Plan
            </button>

            <pre className="bg-gray-100 p-3 rounded">
                {JSON.stringify(tasks, null, 2)}
            </pre>
        </div>
    );
}

export default App;
