import React from 'react';
import { useDrop } from 'react-dnd';
import { ItemTypes } from './ItemTypes';

const LineupSlot = ({ index, player, movePlayerToSlot }) => {
  const [{ isOver }, drop] = useDrop({
    accept: ItemTypes.PLAYER,
    drop: (item) => movePlayerToSlot(item.name, index),
    collect: (monitor) => ({
      isOver: !!monitor.isOver(),
    }),
  });

  return (
    <div ref={drop} className={`lineup-slot ${isOver ? 'is-over' : ''}`}>
      {index + 1}. {player}
    </div>
  );
};

export default LineupSlot;