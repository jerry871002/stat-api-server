import React from 'react';
import { DndProvider } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';
import Lineup from './Lineup';
import Roster from './Roster';

const App = () => {
  const players = [
    'Player 1',
    'Player 2',
    'Player 3',
    'Player 4',
    'Player 5',
    'Player 6',
    'Player 7',
    'Player 8',
    'Player 9',
    'Player 10',
    'Player 11',
    'Player 12',
  ];

  return (
    <DndProvider backend={HTML5Backend}>
      <div className="container">
        <div className="card">
          <Lineup />
        </div>
        <div className="card">
          <Roster players={players} />
        </div>
      </div>
    </DndProvider>
  );
};

export default App;