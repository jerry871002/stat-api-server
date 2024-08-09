import React, { useState } from 'react';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import Lineup from './Lineup';
import Roster from './Roster';
import './App.css';

const App = () => {
  const initialPlayers = [
    { name: 'Mike Trout', avg: 0.304, obp: 0.419, ops: 1.000 },
    { name: 'Mookie Betts', avg: 0.296, obp: 0.374, ops: 0.915 },
    { name: 'Freddie Freeman', avg: 0.295, obp: 0.389, ops: 0.892 },
    { name: 'Juan Soto', avg: 0.287, obp: 0.403, ops: 0.937 },
    { name: 'Fernando Tatis Jr.', avg: 0.282, obp: 0.364, ops: 0.956 },
    { name: 'Ronald Acuña Jr.', avg: 0.281, obp: 0.371, ops: 0.909 },
    { name: 'Nolan Arenado', avg: 0.293, obp: 0.349, ops: 0.890 },
    { name: 'Bryce Harper', avg: 0.276, obp: 0.388, ops: 0.900 },
    { name: 'Jose Altuve', avg: 0.311, obp: 0.360, ops: 0.820 },
    { name: 'Francisco Lindor', avg: 0.285, obp: 0.346, ops: 0.833 },
    { name: 'Christian Yelich', avg: 0.292, obp: 0.381, ops: 0.928 },
    { name: 'Cody Bellinger', avg: 0.273, obp: 0.364, ops: 0.911 },
  ];

  const [players, setPlayers] = useState(initialPlayers.sort((a, b) => a.name.localeCompare(b.name)));
  const [lineup, setLineup] = useState(Array(9).fill(null));
  const [simulationResult, setSimulationResult] = useState(null);

  const movePlayerToSlot = (player, index) => {
    setSimulationResult(null);
    const newLineup = [...lineup];
    const existingPlayer = newLineup[index];
    console.log(existingPlayer);

    // Move the existing player back to the roster if there is one
    if (existingPlayer) {
      setPlayers(prevPlayers => 
        [...prevPlayers, existingPlayer].sort((a, b) => a.name.localeCompare(b.name))
      );
    }

    newLineup[index] = player;
    setLineup(newLineup);
    setPlayers(prevPlayers => 
      prevPlayers.filter(p => p.name !== player.name)
    );
  };

  const removePlayerFromSlot = (index) => {
    setSimulationResult(null);
    const player = lineup[index];
    const newLineup = [...lineup];
    newLineup[index] = null;
    setLineup(newLineup);
    setPlayers([...players, player].sort((a, b) => a.name.localeCompare(b.name)));
  };

  const simulateLineup = async () => {
    // const response = await fetch('https://api.example.com/simulate', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json',
    //   },
    //   body: JSON.stringify({ lineup }),
    // });
    // const result = await response.json();
    const result = { scores: 0.250, hits: 10 };
    setSimulationResult(result);
  };

  return (
    <DndProvider backend={HTML5Backend}>
      <header style={{ textAlign: 'center' }}>
        <h1>Lineup Lab</h1>
      </header>
      <div className="container">
        <div className="card">
          <Lineup lineup={lineup} movePlayerToSlot={movePlayerToSlot} removePlayerFromSlot={removePlayerFromSlot} />
        </div>
        <div className="card">
          <Roster players={players} />
        </div>
      </div>
      <div className="simulate-results">
        <button className="simulate-button" onClick={simulateLineup}>Simulate</button>
        {simulationResult && (
          <div>
            <p>Average Scores: {simulationResult.scores}</p>
            <p>Average Hits: {simulationResult.hits}</p>
          </div>
        )}
      </div>
    </DndProvider>
  );
};

export default App;