import { useState } from "react";

function App() {
    const [prompt, setPrompt] = useState('');
    const [tasks, setTasks] = useState<any[]>([]);

    async function generatePlan() {
        if (!prompt.trim()) return;

        const res = await fetch("http://localhost:8080/agent", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ prompt })
        });

        const data = await res.json();
        setTasks(data);
    }

    return (
        <div className="p-8 max-w-2xl mx-auto space-y-6">
            <h1 className="text-3xl font-bold">Arctiq</h1>

            <textarea
                className="w-full p-3 border rounded"
                rows={4}
                placeholder="What do you want to build?"
                value={prompt}
                onChange={(e) => setPrompt(e.target.value)} />

            <button className="px-4 py-2 bg-blue-600 text-white rounded" onClick={generatePlan}>
                Generate Plan
            </button>

            <div className="space-y-4 mt-6">
                {tasks.map((task, i) => (
                    <div key={i} className="p-4 border rounded shadow-sm bg-white hover:shadow-md transition-shadow">
                        <div className="inline-block px-3 py-1 bg-blue-100 text-blue-700 text-sm font-medium rounded-full">
                            Step {task.Step}
                        </div>
                        <div className="mt-2 text-gray-800 text-lg leading-relaxed">
                            {task.Instruction}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default App;
