package dicts

var en_animals_mammals = []string{
	"Lion", "Tiger", "Elephant", "Leopard", "Cheetah", "Giraffe", "Zebra", "Horse", "Donkey", "Mule", "Camel", "Llama", "Alpaca", "Kangaroo", "Koala", "Wallaby", "Platypus", "Wombat", "Opossum", "Raccoon", "Skunk", "Badger", "Otter", "Weasel", "Mongoose", "Meerkat", "Hyena", "Wolf", "Coyote", "Fox", "Jackal", "Bear", "Panda", "Polar Bear", "Grizzly Bear", "Black Bear", "Brown Bear", "Hippopotamus", "Rhinoceros", "Buffalo", "Bison", "Antelope", "Deer", "Elk", "Moose", "Reindeer", "Caribou", "Sheep", "Goat", "Pig", "Boar", "Warthog", "Hedgehog", "Porcupine", "Beaver", "Squirrel", "Chipmunk", "Mouse", "Rat", "Hamster", "Guinea Pig", "Chinchilla", "Rabbit", "Hare", "Bat", "Whale", "Dolphin", "Porpoise", "Seal", "Sea Lion", "Walrus", "Otter", "Mink", "Ferret", "Lynx", "Bobcat", "Cougar", "Jaguar", "Panther", "Mountain Lion", "Snow Leopard", "Serval", "Ocelot", "Caracal", "Puma", "Manatee", "Dugong", "Sloth", "Armadillo", "Aardvark", "Lemur", "Aye-Aye", "Bushbaby", "Galago", "Tarsier", "Lorikeet",
}
var en_animals_birds = []string{
	"Eagle", "Hawk", "Falcon", "Owl", "Vulture", "Parrot", "Macaw", "Cockatoo", "Canary", "Finch", "Sparrow", "Robin", "Blue Jay", "Cardinal", "Goldfinch", "Oriole", "Hummingbird", "Woodpecker", "Kingfisher", "Heron", "Egret", "Stork", "Flamingo", "Pelican", "Seagull", "Tern", "Puffin", "Albatross", "Penguin", "Swan", "Duck", "Goose", "Mallard", "Teal", "Grebe", "Coot", "Cranes", "Ibis", "Spoonbill", "Kiwi", "Emu", "Ostrich", "Cassowary", "Rhea", "Peacock", "Turkey", "Chicken", "Pheasant", "Quail", "Grouse", "Partridge", "Dove", "Pigeon", "Magpie", "Crow", "Raven", "Jackdaw", "Jay", "Starling", "Blackbird", "Thrush", "Wren", "Nightingale", "Warbler", "Bunting", "Swallow", "Swift", "Lark", "Pipit", "Wagtail", "Flycatcher", "Kinglet", "Tit", "Chickadee", "Nuthatch", "Creeper", "Dipper", "Shrike", "Waxwing", "Sandpiper", "Plovers", "Lapwing", "Curlew", "Whimbrel", "Godwit", "Snipe", "Woodcock", "Phalarope", "Jacana", "Avocet", "Oystercatcher", "Skimmer", "Shearwater", "Petrel", "Fulmar", "Gannet", "Booby", "Cormorant", "Anhinga",
}
var en_animals_fishes = []string{
	"Salmon", "Trout", "Bass", "Cod", "Tuna", "Mackerel", "Herring", "Sardine", "Anchovy", "Catfish", "Carp", "Tilapia", "Perch", "Pike", "Walleye", "Bluegill", "Sunfish", "Crappie", "Sturgeon", "Flounder", "Halibut", "Sole", "Plaice", "Grouper", "Snapper", "Swordfish", "Marlin", "Barracuda", "Shark", "Ray", "Skate", "Eel", "Haddock", "Pollock", "Mahi Mahi", "Sailfish", "Amberjack", "Bonito", "Butterfish", "Char", "Chub", "Drum", "Gar", "Goby", "Grunt", "Hake", "Jack", "Koi", "Lionfish", "Lumpfish", "Menhaden", "Monkfish", "Moray", "Mullet", "Needlefish", "Parrotfish", "Pomfret", "Pompano", "Rabbitfish", "Rockfish", "Sablefish", "Scad", "Scorpionfish", "Sculpin", "Seabass", "Seabream", "Sheepshead", "Silverfish", "Smelt", "Snakehead", "Snapper", "Sprat", "Swordtail", "Tarpon", "Tetra", "Tigerfish", "Trevally", "Triggerfish", "Tripletail", "Trout", "Unicornfish", "Whale Shark", "Wobbegong", "Wrasse", "Yellowtail", "Zander", "Zebrafish", "Flying Fish", "Mudskipper", "Paddlefish", "Pipefish", "Stickleback", "Surgeonfish", "Trumpetfish", "Wormfish", "Anglerfish",
}
var en_animals_insects = []string{
	"Ant", "Bee", "Butterfly", "Moth", "Beetle", "Ladybug", "Firefly", "Dragonfly", "Damselfly", "Grasshopper", "Cricket", "Katydid", "Locust", "Termite", "Cockroach", "Praying Mantis", "Stick Insect", "Leaf Insect", "Cicada", "Aphid", "Leafhopper", "Flea", "Louse", "Mosquito", "Fly", "Horsefly", "Fruit Fly", "Gnat", "Mayfly", "Stonefly", "Caddisfly", "Silverfish", "Firebrat", "Earwig", "Springtail", "Booklouse", "Barklouse", "Thrip", "Biting Midge", "Black Fly", "Deer Fly", "Sandfly", "Blowfly", "Tachinid Fly", "Crane Fly", "Hoverfly", "Robber Fly", "Housefly", "Botfly", "Drosophila", "Weevil", "Stink Bug", "Shield Bug", "Assassin Bug", "Water Strider", "Water Boatman", "Backswimmer", "Giant Water Bug", "Bedbug", "Chinch Bug", "Plant Bug", "Squash Bug", "Kissing Bug", "Lacewing", "Antlion", "Dobsonfly", "Fishfly", "Scorpionfly", "Bark Beetle", "Longhorn Beetle", "Weevil", "Click Beetle", "Bess Beetle", "Stag Beetle", "Hercules Beetle", "Dung Beetle", "Jewel Beetle", "Blister Beetle", "Firefly", "Ladybird Beetle", "Carrion Beetle", "Whirligig Beetle", "Water Beetle", "Ground Beetle", "Darkling Beetle", "Flea Beetle", "Leaf Beetle", "Colorado Potato Beetle", "Tortoise Beetle", "Leafminer", "Sawfly", "Gall Wasp", "Ichneumon Wasp", "Cuckoo Wasp", "Velvet Ant", "Mud Dauber", "Paper Wasp", "Yellowjacket", "Hornet", "Cicada Killer",
}
var en_pets = []string{
	"Dog", "Cat", "Fish", "Bird", "Rabbit", "Hamster", "Guinea Pig", "Turtle", "Ferret", "Lizard", "Snake", "Hedgehog", "Gerbil", "Chinchilla", "Parrot", "Canary", "Lovebird", "Cockatiel", "Budgerigar", "Finch", "Mice", "Rat", "Horse", "Pony", "Pigeon", "Dove", "Frog", "Toad", "Tarantula", "Hermit Crab", "Gecko", "Iguana", "Bearded Dragon", "Tortoise", "Koi Fish", "Betta Fish", "Goldfish", "Axolotl", "Newt", "Salamander", "Sugar Glider", "Capybara", "Mini Pig", "Chickens", "Ducks", "Goats", "Sheep", "Alpacas", "Llamas", "Peafowl", "Quail", "Turkey", "Donkey", "Guinea Fowl", "Cockatoo", "Macaw", "Conure", "Parakeet", "Lovebirds", "Amazon Parrot", "African Grey Parrot", "Eclectus Parrot", "Quaker Parrot", "Ringneck Parrot", "Canaries", "Zebra Finch", "Society Finch", "Java Sparrow", "Diamond Dove", "Pigeon", "Mourning Dove", "Fiddler Crab", "Red Claw Crab", "Blue Crab", "Ghost Shrimp", "Cherry Shrimp", "Bamboo Shrimp", "Snails", "Sea Monkeys", "Clownfish", "Neon Tetra", "Guppy", "Molly", "Platy", "Swordtail", "Angelfish", "Discus Fish", "Corydoras", "Bristlenose Pleco", "Rainbowfish", "African Cichlid", "Oscar Fish", "Gourami", "Ram Cichlid", "Electric Blue Acara", "Kribensis", "Killifish",
}
var en_dog_breeds = []string{
	"Labrador Retriever", "German Shepherd", "Golden Retriever", "French Bulldog", "Bulldog", "Poodle", "Beagle", "Rottweiler", "Yorkshire Terrier", "Boxer", "Dachshund", "Siberian Husky", "Great Dane", "Doberman Pinscher", "Australian Shepherd", "Shih Tzu", "Miniature Schnauzer", "Pembroke Welsh Corgi", "Cavalier King Charles Spaniel", "Shiba Inu", "Chihuahua", "Boston Terrier", "Pomeranian", "Havanese", "Shetland Sheepdog", "Bernese Mountain Dog", "Brittany Spaniel", "English Springer Spaniel", "Cocker Spaniel", "Pug", "Mastiff", "Basset Hound", "Maltese", "Border Collie", "Vizsla", "Newfoundland", "Collie", "Bichon Frise", "Akita", "Weimaraner", "St. Bernard", "West Highland White Terrier", "Rhodesian Ridgeback", "Belgian Malinois", "Bloodhound", "Bull Terrier", "Alaskan Malamute", "Whippet", "Papillon", "Airedale Terrier", "Soft Coated Wheaten Terrier", "Portuguese Water Dog", "Chesapeake Bay Retriever", "Samoyed", "Scottish Terrier", "Staffordshire Bull Terrier", "Cane Corso", "American Bulldog", "Dalmatian", "Italian Greyhound", "Lhasa Apso", "Norwegian Elkhound", "Old English Sheepdog", "Pointer", "Flat-Coated Retriever", "Cairn Terrier", "Basenji", "Anatolian Shepherd", "Afghan Hound", "Chinese Crested", "Schipperke", "Pekingese", "Saluki", "Irish Wolfhound", "Irish Setter", "English Setter", "Gordon Setter", "Coonhound", "Norfolk Terrier", "Norwich Terrier", "Australian Cattle Dog", "American Eskimo Dog", "Japanese Chin", "Tibetan Terrier", "Tibetan Spaniel", "Manchester Terrier", "Rat Terrier", "Keeshond", "Pharaoh Hound", "Treeing Walker Coonhound", "American Hairless Terrier", "German Shorthaired Pointer", "German Wirehaired Pointer", "Clumber Spaniel", "Field Spaniel", "Irish Water Spaniel", "Boykin Spaniel", "Curly-Coated Retriever", "Sussex Spaniel", "Spinone Italiano", "Plott Hound",
}
var en_cat_breeds = []string{
	"Persian", "Maine Coon", "Siamese", "Ragdoll", "British Shorthair", "Sphynx", "Bengal", "Abyssinian", "Birman", "Oriental Shorthair", "Scottish Fold", "Devon Rex", "American Shorthair", "Russian Blue", "Norwegian Forest Cat", "Siberian", "Exotic Shorthair", "Savannah", "Egyptian Mau", "Balinese", "Himalayan", "Chartreux", "Turkish Angora", "Tonkinese", "Burmese", "Cornish Rex", "Japanese Bobtail", "Ocicat", "Somali", "Burmilla", "Ragamuffin", "American Curl", "Turkish Van", "Selkirk Rex", "LaPerm", "Manx", "Snowshoe", "Singapura", "Munchkin", "Toyger", "Bombay", "Nebelung", "Peterbald", "American Bobtail", "Korat", "Havana Brown", "Javanese", "Cymric", "European Shorthair", "Thai", "Donskoy", "German Rex", "Highlander", "Kurilian Bobtail", "Pixie-Bob", "Serengeti", "Sokoke", "Suphalak", "Ukrainian Levkoy", "York Chocolate", "Brazilian Shorthair", "Burmilla Longhair", "Chantilly-Tiffany", "Chausie", "Colorpoint Shorthair", "Cheetoh", "Cyprus", "Dwelf", "Foldex", "German Longhaired", "Khao Manee", "Lykoi", "Mandalay", "Minuet", "Oriental Longhair", "Savannah F1", "Savannah F2", "Savannah F3", "Selkirk Rex Longhair", "Skookum", "Siberian Forest Cat", "Singapura Longhair", "Snow Bengal", "Solid Blue", "Sphynx Longhair", "Thai Lilac", "Thai Blue", "Thai Seal Point", "Thai Chocolate", "Tiffanie", "Tonkinese Longhair", "Toybob", "Ural Rex", "Venetian", "Wegie", "York Chocolate Longhair",
}
var en_pets_names = []string{
	"Max", "Bella", "Charlie", "Luna", "Lucy", "Rocky", "Daisy", "Molly", "Bailey", "Buddy", "Sadie", "Coco", "Lola", "Maggie", "Riley", "Sophie", "Duke", "Zoey", "Bear", "Chloe", "Oliver", "Tucker", "Ruby", "Milo", "Jasper", "Leo", "Oscar", "Harley", "Finn", "Bentley", "Zoe", "Stella", "Toby", "Nala", "Simba", "Penny", "Jake", "Abby", "Rex", "Ginger", "Gizmo", "Jax", "Winston", "Hazel", "Rocco", "Willow", "Lilly", "Louie", "Gracie", "Buster", "Pepper", "Hank", "Ollie", "Ellie", "Sam", "Cleo", "Rudy", "Misty", "Rosie", "Ziggy", "Dexter", "Scout", "Rufus", "Cody", "Bruno", "Hunter", "Mimi", "Smokey", "Frankie", "Mocha", "Midnight", "Boomer", "Ace", "Marley", "Juno", "Chester", "Thor", "Ruby", "Koda", "Benji", "Sasha", "Holly", "Rusty", "Benny", "Maddie", "Shadow", "Murphy", "Piper", "Gus", "Lucky", "Nina", "Blue", "Cookie", "Peanut", "Skye", "Cosmo", "Diesel", "Boo", "Jet",
}
