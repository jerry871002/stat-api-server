import React from 'react';
import { useDrag } from 'react-dnd';
import { ItemTypes } from './ItemTypes';

const Player = ({ name }) => {
  const [{ isDragging }, drag] = useDrag({
    type: ItemTypes.PLAYER,
    item: { name },
    collect: (monitor) => ({
      isDragging: !!monitor.isDragging(),
    }),
  });

  return (
    <div ref={drag} style={{ padding: '10px', border: '1px solid black', marginBottom: '5px', backgroundColor: isDragging ? 'lightblue' : 'white' }}>
      {name}
    </div>
  );
};

export default Player;