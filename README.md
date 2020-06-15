# squirrel
A Raspberry Pi Based Bird Feeder Protector

## A what? 
Squirrels get into bird feeders. It's a fact. So, let's scare them away. `squirrel` is a Raspberry Pi based device that emits a loud 
noise when it dectects the presense of a squirrel. 

## How does it work?
`squirrel` watches the weight of your bird feeder using a load cell. When the weight suddenly changes, it turns on a buzzer. 
The first version of this is using a peizo buzzer at 15 volts which should be emitting around 100db of sound. Hopefully 
it will work. Changes will be made based on learnings. 
