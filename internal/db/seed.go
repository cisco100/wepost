package db

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math/rand"

// 	"github.com/cisco100/wepost/internal/store"
// 	"github.com/google/uuid"
// )

// func SeedDB(store store.Storage) error {
// 	ctx := context.Background()
// 	users := GenerateUSer(500)
// 	posts := GeneratePosts(1000, users)
// 	comments := GenerateComments(1000, users, posts)
// 	for _, user := range users {
// 		if err := store.User.Create(ctx, user); err != nil {
// 			log.Printf("Error creating user %v", err)
// 		}

// 	}

// 	for _, post := range posts {
// 		if err := store.Post.Create(ctx, post); err != nil {
// 			log.Printf("Error creating user %v", err)

// 		}
// 	}

// 	for _, comment := range comments {
// 		if err := store.Comment.Create(ctx, comment); err != nil {
// 			log.Printf("Error creating user %v", err)

// 		}

// 	}

// 	log.Println("Seeding successful!")
// 	return nil
// }

// func GenerateUUID() string {
// 	return uuid.New().String()
// }

// var username = []string{
// 	"User1234", "AlphaWolf", "TechGuru", "BluePhoenix", "ShadowNinja", "CyberKnight", "FireStorm", "QuantumX", "NeonTiger", "DarkVortex",
// 	"SilentHawk", "StormBreaker", "EchoFalcon", "StealthRider", "InfinityEdge", "FrostByte", "SolarFlare", "NightSpecter", "GhostRiderX", "CosmicRay",
// 	"LunarEclipse", "SkyRanger", "TitanForce", "NeuralNet", "AstroVoyager", "HyperDrive", "NeonPulse", "ZeroGravity", "EchoNomad", "WarpSpeed",
// 	"NovaRider", "IceBlade", "ThunderClap", "VortexSurge", "CyberPhantom", "ShadowWalker", "VoidHunter", "InfernoFox", "PhantomStrike", "DarkWolf",
// 	"EchoKnight", "SolarWind", "FrostFire", "QuantumDrift", "NeonSpectra", "TechNomad", "CyberSentinel", "StarlightEcho", "NeuralPulse", "ShadowBlaze",
// 	"HyperNova", "ZeroFusion", "TitanWave", "AstroRogue", "StormChaser", "NovaHunter", "NeonStorm", "DarkHorizon", "WarpTraveler", "QuantumVoid",
// 	"InfernoKnight", "LunarPhoenix", "SkyVanguard", "FrostKnight", "SpectralEcho", "CyberPulse", "SolarKnight", "EchoStorm", "ShadowHorizon", "ZeroChrono",
// 	"PhantomNova", "TitanNomad", "IceVortex", "NeonHorizon", "CosmicEcho", "DarkFury", "QuantumRider", "CyberGhost", "StormVortex", "NeuralDrift",
// 	"FrostNova", "HyperSentinel", "SolarBlaze", "VoidRunner", "EchoStormX", "InfinityHunter", "DarkSpecter", "NeonComet", "AstroKnight", "TitanBlaze",
// 	"ZeroVector", "QuantumBlitz", "SkyBlazer", "PhantomKnight", "WarpShadow", "SolarRogue", "CyberFlame", "HyperStorm", "FrostVortex", "NeonInferno",
// 	"InfernoRider", "StormBreakerX", "TitanFury", "ShadowWarp", "EchoBlitz", "NovaSentinel", "ZeroRogue", "QuantumEdge", "NeuralKnight", "SkyNomad",
// 	"CyberSurge", "SolarEcho", "FrostPhantom", "AstroSpecter", "InfinityVortex", "NeonFlare", "HyperKnight", "DarkPulse", "QuantumStorm", "TitanGhost",
// 	"SolarFusion", "EchoVortex", "ShadowFusion", "VoidRider", "IceStorm", "NovaSpecter", "NeonDrift", "CyberWarp", "LunarSentinel", "SkyDrift",
// 	"FrostEcho", "HyperPhantom", "InfinityDrift", "QuantumNomad", "ShadowComet", "SolarFury", "TitanStorm", "DarkEcho", "NovaWarp", "PhantomBlitz",
// 	"NeonSpecter", "InfernoSentinel", "CyberNova", "WarpFury", "LunarVortex", "SkyFusion", "ZeroHorizon", "QuantumShadow", "NeuralVanguard", "TitanRider",
// 	"EchoFury", "StormRunner", "NeonSurge", "SolarStorm", "HyperEcho", "VoidNomad", "InfinityBlaze", "FrostHunter", "DarkKnightX", "QuantumSpecter",
// 	"CyberStorm", "TitanWarp", "EchoChrono", "NeonBlitz", "LunarFlare", "SkyPhantom", "NovaShadow", "InfernoKnightX", "ShadowCometX", "QuantumHunter",
// 	"ZeroPulse", "SolarSurge", "HyperVortex", "DarkFlare", "StormNomad", "TitanFusion", "NeuralSpecter", "InfinityEdgeX", "WarpKnight", "EchoSentinel",
// 	"CyberFury", "SkyStorm", "NovaFusion", "PhantomDrift", "QuantumWarp", "ShadowPulse", "SolarNova", "DarkSurge", "FrostRogue", "HyperChrono",
// 	"NeonPhantom", "VoidStorm", "TitanEcho", "ZeroFusionX", "QuantumInferno", "SkySpecter", "CyberVanguard", "InfinityComet", "StormHunter", "LunarBlitz",
// 	"DarkRogue", "HyperStormX", "NovaVortex", "SolarWarp", "EchoComet", "ShadowEdge", "ZeroDrift", "QuantumFusion", "CyberPulseX", "TitanNomadX"}

// var title = []string{
// 	"The Future of AI in Healthcare", "Blockchain in Supply Chain Management", "Renewable Energy Innovations", "The Psychology of Marketing", "Advancements in Quantum Computing",
// 	"Global Trade and Economic Policies", "Cybersecurity Trends in 2025", "Agricultural Technologies for Food Security", "Space Exploration and Colonization", "Mental Health Awareness Campaigns",
// 	"The Role of Big Data in Finance", "3D Printing in Manufacturing", "The Rise of Smart Cities", "Human Cloning: Ethical Dilemmas", "Electric Vehicles: The Road Ahead",
// 	"Biotechnology and Genetic Engineering", "Smart Homes and IoT Devices", "The Impact of Social Media on Elections", "Nanotechnology in Medicine", "Green Architecture and Sustainable Living",
// 	"The Evolution of E-commerce", "The Future of Remote Work", "Cultural Heritage Preservation", "The Role of AI in Education", "Cryptocurrency Regulations Worldwide",
// 	"Climate Change and Its Economic Impact", "The Ethics of AI Decision Making", "Wearable Technology and Health Monitoring", "The Future of Virtual Reality", "The Decline of Print Media",
// 	"Autonomous Vehicles and Safety Concerns", "The Influence of Pop Culture on Society", "The Science Behind Happiness", "Advancements in Prosthetic Limbs", "Artificial Intelligence in Art",
// 	"Digital Nomad Lifestyle: Pros and Cons", "The Space Race 2.0: Private Companies", "The Future of Renewable Energy Storage", "Cyberbullying: A Growing Concern", "The Rise of the Metaverse",
// 	"The Impact of 5G on Telecommunications", "The Changing Face of Journalism", "Sports Science and Athletic Performance", "The Role of DNA Testing in Criminal Justice", "The Gig Economy and Worker Rights",
// 	"AI-Powered Personal Assistants", "The Impact of Automation on Jobs", "The Role of Women in STEM", "Smart Farming: The Future of Agriculture", "The Psychology of Consumer Behavior",
// 	"Deepfake Technology and Its Dangers", "The Evolution of Human Rights", "The Role of Government in AI Development", "The Power of Influencer Marketing", "Alternative Medicine: Fact or Fiction?",
// 	"Medical Breakthroughs in Cancer Research", "Ocean Exploration and Its Mysteries", "The Impact of Social Media Algorithms", "The Power of Meditation and Mindfulness", "Global Water Scarcity Solutions",
// 	"Self-Driving Trucks and Logistics", "The Future of Space Tourism", "The Revival of Vinyl Records", "The Dark Web and Cybercrime", "The Science of Sleep",
// 	"Biodegradable Plastics and Their Future", "The Power of Positive Thinking", "The Future of Cryptographic Security", "Solar Energy: Beyond Panels", "AI in Predictive Maintenance",
// 	"Psychological Effects of Virtual Reality", "The Evolution of Online Dating", "The Science of Motivation", "The Influence of Anime on Western Culture", "The Future of Online Education",
// 	"Smart Clothing: The Next Big Thing?", "Ethical Hacking and Cybersecurity", "Crowdfunding and Its Impact on Startups", "The Evolution of Luxury Branding", "The Science of Memory and Learning",
// 	"The Effects of Urbanization on Wildlife", "The Future of Space Mining", "The Role of AI in Music Composition", "The Future of Digital Currencies", "AI in Criminal Investigations",
// 	"Neuralink and Brain-Computer Interfaces", "The Role of Sleep in Cognitive Performance", "Hyperloop: The Future of Transport?", "Climate Change and Rising Sea Levels", "The Art of Minimalist Living",
// 	"The Power of Storytelling in Business", "The Evolution of Video Games", "The Role of Robotics in Disaster Relief", "AI and the Future of Customer Service", "The Science of Aging",
// 	"The Rise of Subscription-Based Services", "The History and Future of Tattoos", "The Power of Body Language in Communication", "The Ethics of Genetic Modification", "The Influence of K-Pop Globally",
// 	"Social Media Addiction and Its Effects", "The Science of Attraction", "Renewable Energy Jobs and the Future Workforce", "The Evolution of Cyber Warfare", "The Future of Global Internet Access",
// 	"The Role of AI in Predicting Natural Disasters", "The Psychology of Procrastination", "The Future of Autonomous Drones", "The Evolution of Space Suits", "The Power of Branding in Business",
// 	"Edible Packaging and Sustainability", "The Science of Humor", "AI and Personalized Medicine", "The Ethics of Surveillance Technology", "Virtual Reality Therapy for PTSD",
// 	"The Role of AI in Wildlife Conservation", "The Future of Personalized Advertising", "The Evolution of Street Fashion", "The Influence of Science Fiction on Innovation", "The Power of Crowdsourced Data",
// 	"The Rise of Biometric Authentication", "The Science of Lucid Dreaming", "The Future of Hyper-Personalized Learning", "The Role of AI in Detecting Fake News", "The Evolution of Digital Payments",
// 	"The Power of Color Psychology", "AI-Powered Stock Trading", "The Science of Deja Vu", "The Role of AI in Enhancing Creativity", "The Influence of Music on Productivity",
// 	"Human Augmentation and the Future of Work", "Sustainable Fashion and Eco-Friendly Clothing", "The Evolution of Handwriting in the Digital Age", "AI-Powered Fraud Detection", "The Role of Virtual Influencers",
// 	"3D Printed Food and the Future of Dining", "The Science of Crowd Psychology", "The Rise of Ethical Consumerism", "The Future of Global Banking", "AI and Predictive Policing",
// 	"The Role of Smart Contracts in Business", "The Evolution of Fantasy Sports", "Personalized AI Nutrition Plans", "The Future of Online Anonymity", "The Power of Nostalgia Marketing",
// 	"The Ethics of AI-Generated Content", "The Influence of Esports on Traditional Sports", "The Science Behind Laughter", "The Evolution of Fast Food Industry", "The Role of AI in Personalized Coaching",
// 	"AI-Powered Chatbots in Healthcare", "The Science of Dopamine and Motivation", "The Future of Microtransactions in Gaming", "AI in Urban Planning", "The Power of Habits in Daily Life",
// 	"The Evolution of Home Automation", "The Influence of Digital Art on Traditional Art", "AI and the Future of Restaurant Automation", "The Role of AI in Resume Screening", "The Science of Decision Making",
// 	"The Psychology of Viral Content", "The Future of AI-Generated Music", "The Role of AI in Legal Research", "AI-Powered Mental Health Support", "The Evolution of Public Speaking Techniques",
// 	"The Influence of Science on Fiction Writing", "The Science of Sensory Perception", "The Role of AI in Archaeological Discoveries", "AI in Video Surveillance and Security", "The Future of Nanomedicine",
// 	"The Influence of Art on Mental Wellbeing", "The Role of AI in Language Translation", "AI-Powered Virtual Interior Design", "The Science of Time Perception", "AI in Personalized News Feeds",
// 	"The Future of AI-Powered Tutors", "The Ethics of AI in Political Campaigns", "The Influence of Space Technology on Daily Life", "AI in Personalized Job Recommendations", "The Science of Brain Plasticity",
// 	"The Future of AI-Powered Diagnostics", "The Role of AI in Disaster Management", "The Science of Smell and Memory", "The Rise of Fintech Startups", "AI in Personalized Healthcare",
// 	"The Psychology of Brand Loyalty", "The Evolution of Online Marketplaces", "AI in Space Exploration", "The Role of 3D Printing in Aerospace", "AI and the Future of Human Creativity",
// 	"The Science of Habit Formation", "The Ethics of AI-Generated Art", "Smart Roads and Traffic Management", "The Rise of Virtual Reality Gaming", "The Influence of Genetics on Personality",
// 	"The Role of AI in Financial Forecasting", "The Future of Quantum Cryptography", "AI and the Future of Human-Computer Interaction", "The Impact of Microplastics on Marine Life", "The Future of Renewable Hydrogen",
// 	"The Role of AI in Personalized Shopping", "AI in Smart Factories", "The Science of Productivity Hacks", "The Future of Personalized Learning Apps", "The Rise of Decentralized Social Media",
// 	"The Future of Precision Agriculture", "AI and the Evolution of Digital Assistants", "The Impact of Automation on the Creative Industry", "The Psychology of Online Shopping", "The Science of Muscle Memory",
// 	"The Future of Biometric Payments", "The Role of AI in Smart Grid Optimization", "The Ethics of AI-Powered Surveillance", "The Science of Emotional Intelligence", "The Evolution of Cloud Computing",
// 	"AI and the Future of Supply Chain Management", "The Rise of AI-Generated Content", "The Science of Sound and Healing", "The Future of Personalized Workout Plans", "The Role of AI in Real Estate Valuation",
// 	"The Future of Blockchain in Voting Systems", "The Science of Cold Exposure Therapy", "AI and the Future of Cyber Defense", "The Evolution of Subscription Services", "The Impact of AI on Video Production",
// 	"The Role of AI in Personalized Advertising", "The Science of Time Management", "The Evolution of Ride-Sharing Services", "The Future of AI-Generated Fashion", "The Ethics of AI in Hiring",
// 	"The Psychology of Music Therapy", "The Future of Personalized Virtual Reality", "The Science of Human-Computer Interaction", "The Role of AI in Customer Retention", "The Future of Augmented Reality Retail",
// 	"AI in Personalized Therapy", "The Evolution of Smart Textiles", "The Science of Personalized Nutrition", "The Role of AI in Wildlife Protection", "The Future of AI-Generated Literature",
// 	"The Impact of AI on Small Businesses", "The Evolution of Smart Assistants", "The Science of Decision Fatigue", "The Role of AI in Predicting Climate Change", "The Future of AI in Journalism",
// 	"The Ethics of AI-Powered Medical Research", "The Science of Psychological Resilience", "The Role of AI in Disaster Relief Coordination", "The Future of AI in Personalized Finance", "The Evolution of AI-Generated Influencers",
// 	"The Role of AI in Detecting Online Misinformation", "The Science of Neuroplasticity", "The Future of AI in Employee Training", "The Impact of AI on Customer Experience", "The Role of AI in Home Security",
// 	"The Future of AI-Powered Chatbots", "The Evolution of AI in Music Production", "The Science of Human Behavior Prediction", "The Role of AI in Smart Home Automation", "The Future of AI in Environmental Monitoring",
// 	"The Science of Sleep Optimization", "The Evolution of AI in Retail Analytics", "The Future of AI-Powered Traffic Management", "The Role of AI in Remote Patient Monitoring", "The Science of AI-Powered Personalized Medicine",
// 	"The Impact of AI on the Stock Market", "The Role of AI in AI-Powered Customer Service", "The Future of AI in Content Moderation", "The Science of AI-Powered Creativity", "The Role of AI in Predicting Financial Trends",
// 	"The Future of AI-Powered Business Strategy", "The Science of AI-Powered Personalization", "The Role of AI in Smart Cities Development", "The Future of AI in Predicting Mental Health Issues", "The Evolution of AI in Sports Analytics",
// 	"The Future of AI in Emergency Response", "The Role of AI in AI-Powered HR Management", "The Science of AI-Powered Behavioral Analysis", "The Future of AI in Public Safety", "The Role of AI in AI-Powered Autonomous Vehicles",
// 	"The Evolution of AI in AI-Powered Virtual Reality", "The Science of AI-Powered Emotional Recognition", "The Future of AI in AI-Powered Smart Homes", "The Role of AI in AI-Powered Predictive Maintenance", "The Future of AI in AI-Powered Fraud Detection",
// 	"The Evolution of AI in AI-Powered Personalized Education", "The Science of AI-Powered Brain-Computer Interfaces", "The Role of AI in AI-Powered Personalized Learning", "The Future of AI in AI-Powered Personalized Healthcare", "The Science of AI-Powered Robotics",
// 	"The Evolution of AI in AI-Powered Personalized Marketing", "The Role of AI in AI-Powered Personalized Security", "The Future of AI in AI-Powered Personalized Shopping", "The Science of AI-Powered Personalized Transportation", "The Role of AI in AI-Powered Personalized Entertainment",
// 	"The Evolution of AI in AI-Powered Personalized Sports Analytics", "The Science of AI-Powered Personalized Autonomous Vehicles", "The Future of AI in AI-Powered Personalized Smart Home Automation", "The Role of AI in AI-Powered Personalized Workplace Optimization", "The Science of AI-Powered Personalized Virtual Assistants",
// 	"The Future of AI in AI-Powered Personalized Customer Support", "The Role of AI in AI-Powered Personalized Cybersecurity", "The Evolution of AI in AI-Powered Personalized Public Safety", "The Science of AI-Powered Personalized Financial Planning", "The Future of AI in AI-Powered Personalized Political Campaigning",
// 	"The Role of AI in AI-Powered Personalized Energy Management", "The Science of AI-Powered Personalized Manufacturing", "The Evolution of AI in AI-Powered Personalized Aerospace Engineering", "The Future of AI in AI-Powered Personalized Environmental Conservation", "The Role of AI in AI-Powered Personalized Public Transport",
// 	"The Science of AI-Powered Personalized Retail Shopping", "The Evolution of AI in AI-Powered Personalized Sports Training", "The Future of AI in AI-Powered Personalized Legal Assistance", "The Role of AI in AI-Powered Personalized Virtual Worlds", "The Science of AI-Powered Personalized Hospitality",
// 	"The Future of AI in AI-Powered Personalized Healthcare Research", "The Role of AI in AI-Powered Personalized Digital Marketing", "The Evolution of AI in AI-Powered Personalized Scientific Research", "The Science of AI-Powered Personalized Space Exploration", "The Future of AI in AI-Powered Personalized Weather Prediction"}

// var content = []string{"Explores how AI is revolutionizing medical diagnostics and patient care.",
// 	"Discusses AI’s role in predicting and managing natural disasters.",
// 	"Examines the link between smells and human memory retention.",
// 	"Analyzes the rapid rise and impact of fintech startups.",
// 	"Explores how AI personalizes healthcare treatment and diagnostics.",

// 	"Breaks down the psychology behind brand loyalty and consumer behavior.",
// 	"Explores the evolution of online marketplaces and e-commerce trends.",
// 	"Examines AI’s contributions to space exploration and research.",
// 	"Discusses the role of 3D printing in aerospace innovation.",
// 	"Analyzes AI’s influence on creativity and artistic expression.",

// 	"Explores how habits are formed and how they can be changed.",
// 	"Examines the ethical implications of AI-generated artwork.",
// 	"Discusses the future of smart roads and AI-driven traffic systems.",
// 	"Analyzes the rise of virtual reality gaming and its impact on entertainment.",
// 	"Examines how genetics influence human personality traits.",

// 	"Explores how AI is used for financial forecasting and market predictions.",
// 	"Discusses the future of quantum cryptography in cybersecurity.",
// 	"Analyzes AI’s impact on human-computer interactions.",
// 	"Explores the dangers of microplastics in marine ecosystems.",
// 	"Discusses hydrogen’s role as a sustainable energy source.",

// 	"Examines how AI powers personalized online shopping experiences.",
// 	"Explores how AI optimizes smart factories for efficiency.",
// 	"Discusses the science behind productivity hacks and time management.",
// 	"Analyzes the impact of AI in personalized e-learning applications.",
// 	"Explores decentralized social media platforms and data privacy concerns.",

// 	"Discusses AI-driven precision agriculture for better crop management.",
// 	"Analyzes the evolution of digital assistants like Alexa and Siri.",
// 	"Explores how automation is reshaping the creative industries.",
// 	"Examines the psychology behind impulse buying and online shopping.",
// 	"Discusses the science behind muscle memory and skill retention.",

// 	"Explores how biometric payments are revolutionizing financial transactions.",
// 	"Examines AI’s role in smart grid optimization and energy efficiency.",
// 	"Discusses the ethical concerns of AI-powered surveillance systems.",
// 	"Explores the science behind emotional intelligence and human behavior.",
// 	"Analyzes the evolution of cloud computing and its impact on businesses.",

// 	"Examines AI’s role in supply chain management and logistics.",
// 	"Explores the rise of AI-generated content in digital media.",
// 	"Discusses how sound therapy aids in mental and physical healing.",
// 	"Explores how AI personalizes fitness routines and workout plans.",
// 	"Examines how AI improves real estate valuation and market predictions.",

// 	"Discusses how blockchain can enhance voting system security.",
// 	"Explores the science behind cold exposure therapy and its benefits.",
// 	"Examines AI’s impact on cybersecurity and defense strategies.",
// 	"Discusses the rise of subscription-based business models.",
// 	"Analyzes how AI is transforming video production and editing.",

// 	"Explores AI-driven personalized advertising and marketing strategies.",
// 	"Discusses effective time management techniques for professionals.",
// 	"Analyzes the evolution of ride-sharing services like Uber and Lyft.",
// 	"Explores AI’s role in fashion design and personalized styling.",
// 	"Examines the ethics of AI in hiring and recruitment processes.",

// 	"Discusses how music therapy benefits mental health and well-being.",
// 	"Explores the future of personalized virtual reality experiences.",
// 	"Examines how AI enhances human-computer interaction.",
// 	"Discusses AI’s role in customer retention and engagement strategies.",
// 	"Explores the use of augmented reality in retail shopping experiences.",

// 	"Examines AI-driven therapy solutions for mental health treatment.",
// 	"Discusses innovations in smart textiles and wearable technology.",
// 	"Explores the science behind personalized nutrition and diet plans.",
// 	"Examines AI’s role in wildlife conservation and protection efforts.",
// 	"Discusses the rise of AI-generated literature and storytelling.",

// 	"Analyzes how AI is benefiting small businesses and entrepreneurs.",
// 	"Explores the advancements in AI-powered voice assistants.",
// 	"Discusses the science behind decision fatigue and cognitive overload.",
// 	"Examines AI’s role in predicting climate change and its effects.",
// 	"Discusses how AI is shaping the future of journalism.",

// 	"Explores the ethical concerns of AI in medical research.",
// 	"Examines the science behind psychological resilience and adaptability.",
// 	"Discusses AI’s role in disaster relief coordination and response.",
// 	"Explores AI-driven financial planning and investment strategies.",
// 	"Examines the rise of AI-generated influencers in social media.",

// 	"Discusses AI’s role in detecting and preventing online misinformation.",
// 	"Explores the science behind neuroplasticity and brain adaptation.",
// 	"Examines AI-driven employee training and workplace learning.",
// 	"Discusses AI’s impact on customer experience and brand interaction.",
// 	"Explores AI-powered home security systems and smart surveillance.",

// 	"Examines AI’s role in chatbot development and conversational AI.",
// 	"Discusses how AI is revolutionizing music production and composition.",
// 	"Explores how AI predicts human behavior through data analysis.",
// 	"Examines AI-driven automation in smart home technology.",
// 	"Discusses AI’s role in environmental monitoring and conservation.",

// 	"Analyzes the science behind sleep optimization and better rest.",
// 	"Examines AI’s impact on retail analytics and customer behavior.",
// 	"Explores AI-powered traffic management systems for cities.",
// 	"Discusses AI’s role in remote patient monitoring and telemedicine.",
// 	"Examines AI’s contributions to personalized medicine and healthcare.",

// 	"Analyzes AI’s influence on stock market predictions and trading.",
// 	"Explores how AI improves customer service and support systems.",
// 	"Discusses AI-driven moderation of digital content and social media.",
// 	"Examines AI’s role in creative fields like design and media production.",
// 	"Explores how AI predicts financial trends and market fluctuations.",

// 	"Discusses AI-powered business strategies for decision-making.",
// 	"Examines the impact of AI in personalized advertising and marketing.",
// 	"Explores AI’s role in the development of smart cities.",
// 	"Discusses AI-driven mental health detection and intervention.",
// 	"Examines AI-powered analytics in professional sports.",

// 	"Explores how AI assists in emergency response and crisis management.",
// 	"Discusses AI’s role in HR management and employee recruitment.",
// 	"Examines AI-driven behavioral analysis for security and law enforcement.",
// 	"Explores AI-powered public safety and disaster preparedness.",
// 	"Discusses AI-driven development of autonomous vehicles.",

// 	"Analyzes AI’s role in virtual reality innovation and experiences.",
// 	"Explores AI-powered emotional recognition and human interactions.",
// 	"Discusses AI-driven automation in smart homes and IoT devices.",
// 	"Examines AI’s impact on predictive maintenance in industries.",
// 	"Explores AI-powered fraud detection in financial transactions.",

// 	"Analyzes AI’s role in education and personalized learning tools.",
// 	"Discusses AI-powered brain-computer interfaces and neuroscience.",
// 	"Explores AI-driven personalized learning for students.",
// 	"Examines AI-powered precision healthcare and treatment plans.",
// 	"Discusses AI’s impact on robotics and autonomous systems.",
// 	"Explores how AI is used in predicting political election outcomes.",
// 	"Examines the role of AI in personalized job recommendations.",
// 	"Discusses AI-driven advancements in drug discovery and development.",
// 	"Analyzes the future of space tourism and AI-powered astronaut training.",
// 	"Explores the impact of AI in personalized shopping recommendations.",

// 	"Discusses AI’s role in fighting climate change and reducing carbon footprints.",
// 	"Examines the role of AI in predictive maintenance for industrial equipment.",
// 	"Analyzes how AI is enhancing cybersecurity and data protection.",
// 	"Explores the integration of AI in smart agriculture and crop monitoring.",
// 	"Discusses AI-powered fraud detection in online banking and finance.",

// 	"Explores AI-driven content generation in digital marketing and SEO.",
// 	"Analyzes the rise of AI in real-time language translation.",
// 	"Discusses AI’s role in automating legal document review and analysis.",
// 	"Examines how AI is transforming logistics and supply chain efficiency.",
// 	"Explores AI-driven personalized mental health therapy.",

// 	"Discusses how AI is optimizing renewable energy management.",
// 	"Analyzes AI’s impact on movie and TV show recommendations.",
// 	"Explores AI-powered automation in manufacturing industries.",
// 	"Discusses AI-driven predictive analytics in sports coaching.",
// 	"Examines AI’s role in disaster risk assessment and response planning.",

// 	"Explores how AI is used in developing self-driving public transport.",
// 	"Discusses AI’s contributions to smart wearable technology.",
// 	"Analyzes AI-powered job recruitment and candidate screening.",
// 	"Explores AI-driven sentiment analysis in social media monitoring.",
// 	"Discusses AI’s role in automated customer service chatbots.",

// 	"Examines AI-powered automation in warehouse and logistics management.",
// 	"Explores AI-driven fraud detection in e-commerce transactions.",
// 	"Discusses how AI is used to enhance online education experiences.",
// 	"Analyzes AI’s impact on predicting global economic trends.",
// 	"Explores AI-driven personalized fitness coaching applications.",

// 	"Discusses how AI is revolutionizing robotic-assisted surgeries.",
// 	"Examines AI-powered translation tools for real-time multilingual communication.",
// 	"Explores AI-driven trend forecasting in the fashion industry.",
// 	"Discusses AI’s role in optimizing digital advertising campaigns.",
// 	"Analyzes AI-powered speech recognition advancements.",

// 	"Explores how AI is used to create deepfake detection technology.",
// 	"Discusses AI-driven personalization in news and media consumption.",
// 	"Examines AI-powered automation in real estate property valuation.",
// 	"Explores AI-driven enhancements in augmented reality experiences.",
// 	"Discusses AI-powered advancements in autonomous drone technology.",

// 	"Analyzes AI’s role in social media content moderation.",
// 	"Explores AI-driven automation in smart home security systems.",
// 	"Discusses AI-powered financial risk analysis for investments.",
// 	"Examines AI-driven voice cloning and its ethical implications.",
// 	"Explores how AI is used in facial recognition security systems.",

// 	"Discusses AI-driven efficiency improvements in hospital management.",
// 	"Analyzes AI’s role in developing smart assistants for businesses.",
// 	"Explores AI-powered automation in stock market trading.",
// 	"Discusses AI-driven innovations in self-checkout and cashier-less stores.",
// 	"Examines AI-powered handwriting recognition technology.",

// 	"Explores AI-driven automation in public transportation scheduling.",
// 	"Discusses AI-powered automation in call center operations.",
// 	"Analyzes AI’s impact on real-time disease outbreak prediction.",
// 	"Explores AI-driven traffic flow optimization in urban planning.",
// 	"Discusses AI-powered automation in fast food and restaurant services.",

// 	"Examines AI’s role in predictive policing and crime prevention.",
// 	"Explores AI-driven sentiment analysis in financial markets.",
// 	"Discusses AI-powered automation in event planning and management.",
// 	"Analyzes AI-driven efficiency improvements in airport security.",
// 	"Explores AI-powered analytics for personalized book recommendations.",

// 	"Discusses AI-driven innovations in sports commentary and analysis.",
// 	"Examines AI-powered predictive analytics in insurance risk assessment.",
// 	"Explores AI-driven improvements in personalized travel experiences.",
// 	"Discusses AI-powered automation in healthcare appointment scheduling.",
// 	"Analyzes AI-driven automation in online legal consultation.",

// 	"Explores AI-powered solutions for optimizing water resource management.",
// 	"Discusses AI-driven automation in personalized interior design.",
// 	"Examines AI’s role in optimizing road construction and traffic planning.",
// 	"Analyzes AI-powered automation in smart waste management systems.",
// 	"Explores AI-driven automation in wildlife conservation monitoring.",

// 	"Discusses AI-powered automation in city-wide energy efficiency management.",
// 	"Examines AI-driven automation in personalized meal planning and cooking.",
// 	"Explores AI’s role in personalized customer service interactions.",
// 	"Discusses AI-driven automation in transportation route optimization.",
// 	"Analyzes AI-powered solutions for smart school campus management.",

// 	"Explores AI-powered automation in water desalination and purification.",
// 	"Discusses AI-driven automation in smart farming equipment.",
// 	"Examines AI’s impact on reducing traffic congestion in cities.",
// 	"Analyzes AI-powered real-time weather forecasting improvements.",
// 	"Explores AI-driven automation in news article summarization.",
// 	"Discusses AI-powered automation in optimizing marine navigation routes.",
// 	"Explores AI-driven automation in detecting and preventing tax fraud.",
// 	"Analyzes AI’s role in optimizing HVAC systems for energy efficiency.",
// 	"Examines AI-powered automation in designing next-gen aircraft.",
// 	"Discusses AI’s impact on automating corporate financial forecasting.",

// 	"Explores AI-driven automation in tracking deforestation and illegal logging.",
// 	"Analyzes AI-powered automation in translating ancient manuscripts.",
// 	"Discusses AI’s role in creating immersive virtual museum tours.",
// 	"Examines AI-driven automation in preventing traffic accidents.",
// 	"Explores AI-powered automation in creating realistic virtual influencers.",

// 	"Discusses AI’s impact on assisting students with personalized study plans.",
// 	"Analyzes AI-driven automation in forecasting power grid failures.",
// 	"Explores AI-powered automation in optimizing remote work productivity.",
// 	"Examines AI’s role in streamlining movie scriptwriting processes.",
// 	"Discusses AI-driven automation in enabling real-time language dubbing.",

// 	"Explores AI-powered automation in assisting chefs with recipe innovations.",
// 	"Analyzes AI’s impact on automating toxic waste management.",
// 	"Discusses AI-driven automation in helping users manage digital subscriptions.",
// 	"Examines AI-powered automation in identifying deepfake images and videos.",
// 	"Explores AI’s role in improving crowd management at large events.",

// 	"Discusses AI-driven automation in securing online voting systems.",
// 	"Analyzes AI-powered automation in automating drone-assisted search missions.",
// 	"Explores AI’s role in personalized skincare formulations using AI.",
// 	"Examines AI-powered automation in detecting construction safety hazards.",
// 	"Discusses AI-driven automation in optimizing airport baggage handling.",

// 	"Explores AI’s impact on enhancing voice recognition for disabled users.",
// 	"Analyzes AI-powered automation in monitoring industrial air pollution.",
// 	"Discusses AI-driven automation in tracking and preserving endangered languages.",
// 	"Examines AI-powered automation in accelerating vaccine development.",
// 	"Explores AI’s role in optimizing maritime cargo transport logistics.",

// 	"Discusses AI-powered automation in creating ultra-realistic video game graphics.",
// 	"Analyzes AI-driven automation in monitoring blockchain transactions.",
// 	"Explores AI’s role in real-time translation for international business negotiations.",
// 	"Examines AI-powered automation in preventing online misinformation.",
// 	"Discusses AI-driven automation in designing sustainable smart buildings.",

// 	"Explores AI’s impact on enhancing robotic-assisted farming techniques.",
// 	"Analyzes AI-powered automation in predicting and preventing software bugs.",
// 	"Discusses AI-driven automation in optimizing wind farm energy output.",
// 	"Examines AI’s role in providing real-time automated weather alerts.",
// 	"Explores AI-powered automation in digitizing and restoring historical photos.",

// 	"Discusses AI-driven automation in personalized career path recommendations.",
// 	"Analyzes AI’s role in automating cybersecurity threat detection.",
// 	"Explores AI-powered automation in streamlining hospital bed management.",
// 	"Examines AI-driven automation in optimizing solar panel efficiency.",
// 	"Discusses AI’s role in automating precision-guided medical treatments.",

// 	"Explores AI-powered automation in real-time wildfire detection and response.",
// 	"Analyzes AI-driven automation in automating urban waste collection.",
// 	"Discusses AI’s role in providing hyper-personalized health and fitness plans.",
// 	"Examines AI-powered automation in managing real-time crowd movements.",
// 	"Explores AI-driven automation in optimizing space-based telescope operations.",

// 	"Discusses AI’s impact on personalizing financial planning for individuals.",
// 	"Analyzes AI-powered automation in streamlining restaurant order processing.",
// 	"Explores AI’s role in optimizing smart classroom learning experiences.",
// 	"Examines AI-driven automation in designing AI-generated interior design themes.",
// 	"Discusses AI-powered automation in detecting counterfeit goods and currency.",

// 	"Explores AI’s impact on enabling self-sustaining smart cities.",
// 	"Analyzes AI-powered automation in helping automate self-driving ship navigation.",
// 	"Discusses AI-driven automation in optimizing industrial chemical processes.",
// 	"Examines AI’s role in developing AI-assisted customer experience agents.",
// 	"Explores AI-powered automation in managing automated smart parking systems.",

// 	"Discusses AI-driven automation in predicting infrastructure maintenance needs.",
// 	"Analyzes AI-powered automation in optimizing personalized home automation.",
// 	"Explores AI’s role in ensuring supply chain traceability and transparency.",
// 	"Examines AI-driven automation in streamlining electronic waste recycling.",
// 	"Discusses AI-powered automation in tracking food nutrition labeling accuracy.",

// 	"Explores AI’s impact on improving digital twin simulations for urban planning.",
// 	"Analyzes AI-driven automation in automating customer complaint resolutions.",
// 	"Discusses AI-powered automation in streamlining academic paper review.",
// 	"Examines AI’s role in automating AI-powered pharmaceutical research.",
// 	"Explores AI-driven automation in optimizing drone-assisted package deliveries.",

// 	"Discusses AI-powered automation in preventing illegal wildlife trade.",
// 	"Analyzes AI’s impact on automating construction blueprint analysis.",
// 	"Explores AI-driven automation in streamlining microchip manufacturing.",
// 	"Examines AI-powered automation in enhancing AI-driven human emotion analysis.",
// 	"Discusses AI’s role in detecting anomalies in space-based satellite imagery.",

// 	"Explores AI-powered automation in assisting with AI-generated art creation.",
// 	"Analyzes AI’s impact on automating DNA analysis in forensic investigations.",
// 	"Discusses AI-driven automation in enabling smart city traffic analytics.",
// 	"Examines AI-powered automation in personalizing meditation and relaxation apps.",
// 	"Explores AI’s role in optimizing real-time emergency disaster logistics.",

// 	"Discusses AI-powered automation in ensuring AI-powered healthcare fraud detection.",
// 	"Analyzes AI’s impact on automating self-learning AI tutoring bots.",
// 	"Explores AI-driven automation in monitoring global economic fluctuations.",
// 	"Examines AI-powered automation in optimizing car insurance premium calculations.",
// 	"Discusses AI’s role in streamlining automated handwriting-to-text conversion.",

// 	"Explores AI-powered automation in enhancing quantum computing error correction.",
// 	"Analyzes AI-driven automation in predicting sports match outcomes.",
// 	"Discusses AI-powered automation in automating AI-generated voiceovers.",
// 	"Examines AI’s role in optimizing network traffic management.",
// 	"Explores AI-driven automation in preventing illegal deep-sea fishing activities.",

// 	"Discusses AI-powered automation in streamlining the process of digital identity verification.",
// 	"Analyzes AI’s impact on automating customer loyalty program personalization.",
// 	"Explores AI-driven automation in detecting fake news and disinformation campaigns.",
// 	"Examines AI-powered automation in assisting with automated legal dispute resolution.",
// 	"Discusses AI’s role in streamlining AI-powered stock market analytics.",
// }

// var comment = []string{"The future of AI in healthcare looks promising.", "Renewable energy is the key to sustainability.", "Space tourism might be mainstream in a decade.", "Cryptocurrency is reshaping global finance.", "Self-driving cars will change urban landscapes.",
// 	"Cybersecurity threats are evolving rapidly.", "3D printing is revolutionizing manufacturing.", "Online education is making learning accessible.", "Agriculture is embracing automation.", "The metaverse could redefine social interaction.",
// 	"Electric vehicles are gaining popularity.", "Climate change is the biggest challenge we face.", "The rise of biotech will extend human lifespan.", "Fashion trends are going circular with recycling.", "Drones are transforming logistics.",
// 	"Quantum computing will unlock new possibilities.", "Smart cities will enhance urban living.", "Blockchain is revolutionizing supply chains.", "AI-generated content is becoming indistinguishable.", "Wearable tech is improving health monitoring.",
// 	"E-sports is a billion-dollar industry now.", "Neural interfaces could merge humans with machines.", "The gig economy is reshaping employment.", "Remote work is the new normal.", "Medical advancements are curing once-fatal diseases.",
// 	"The real estate market is embracing digital platforms.", "Personalized medicine is improving treatment outcomes.", "Autonomous robots are assisting in surgeries.", "AI ethics debates are heating up.", "Augmented reality is enhancing retail experiences.",
// 	"Biodegradable packaging is reducing waste.", "Renewable energy investments are increasing.", "Social media influencers are shaping brands.", "Voice assistants are making tech more intuitive.", "5G networks are enabling real-time interactions.",
// 	"Ocean exploration is revealing hidden ecosystems.", "Gene editing is pushing ethical boundaries.", "SpaceX is making Mars colonization possible.", "Gaming industry revenues rival Hollywood.", "Traditional banking is being disrupted by fintech.",
// 	"AI-powered chatbots are improving customer service.", "Vertical farming is solving urban food shortages.", "Renewable hydrogen is a future fuel source.", "Smart contracts are transforming legal agreements.", "Tech startups are driving innovation globally.",
// 	"Mental health awareness is growing worldwide.", "Cryptographic security is crucial in digital transactions.", "The circular economy is minimizing resource waste.", "DNA storage might be the future of data preservation.", "Hyperloop technology could revolutionize travel."}

// var tags = []string{
// 	"AI-Diagnostics", "Disaster-AI", "Memory-Science", "Fintech", "Healthcare-AI",
// 	"Brand-Loyalty", "E-Commerce", "Space-AI", "Aerospace-3D", "AI-Creativity",
// 	"Habit-Science", "AI-Art", "Smart-Traffic", "VR-Gaming", "Genetics-Personality",
// 	"Finance-AI", "Quantum-Crypto", "HCI-AI", "Microplastics", "Hydrogen-Energy",
// 	"AI-Shopping", "Smart-Factories", "Productivity-Hacks", "E-Learning", "Decentralized-Social",
// 	"Precision-Ag", "Digital-Assistants", "Automation-Creativity", "E-Commerce-Psych", "Muscle-Memory",
// 	"Biometric-Payments", "Smart-Grid", "AI-Surveillance", "Emotional-Intelligence", "Cloud-Computing",
// 	"AI-Supply-Chain", "AI-Content", "Sound-Healing", "Fitness-Tech", "Real-Estate-AI",
// 	"Blockchain-Voting", "Cold-Therapy", "Cyber-Defense", "Subscription-Economy", "AI-Video",
// 	"AI-Advertising", "Time-Management", "Ride-Sharing", "Fashion-AI", "HR-AI",
// 	"Music-Therapy", "VR-Personalization", "UX-Science", "Customer-Retention", "AR-Retail",
// 	"Therapy-AI", "Smart-Textiles", "Personalized-Nutrition", "Wildlife-AI", "AI-Literature",
// 	"SME-AI", "Voice-Assistants", "Decision-Fatigue", "Climate-Predict", "Journalism-AI",
// 	"Medical-Ethics", "Resilience-Science", "Disaster-AI", "Finance-AI", "Influencer-AI",
// 	"Misinformation-AI", "Neuroplasticity", "Corporate-Training", "Customer-Experience", "Home-Security",
// 	"Chatbots", "AI-Music", "Behavior-Prediction", "Smart-Homes", "Eco-AI",
// 	"Sleep-Tech", "Retail-AI", "Traffic-AI", "Telemedicine", "Personalized-Medicine",
// 	"Stock-AI", "Customer-AI", "Content-AI", "Creative-AI", "Finance-Predict",
// 	"Business-AI", "Personalization-AI", "Smart-Cities", "Mental-Health-AI", "Sports-Analytics",
// 	"Emergency-AI", "HR-Tech", "Behavior-AI", "Public-Safety", "Autonomous-AI",
// 	"VR-AI", "Emotion-AI", "Smart-Homes-AI", "Maintenance-AI", "Fraud-AI",
// 	"EdTech-AI", "Brain-Computer", "E-Learning-AI", "Healthcare-AI", "Robotics-AI",
// 	"Marketing-AI", "Security-AI", "E-Commerce-AI", "Transportation-AI", "Entertainment-AI",
// 	"SportsTech-AI", "Autonomous-Vehicles", "Smart-Home-AI", "Workplace-AI", "Virtual-Assistants",
// 	"Customer-Support", "Cybersecurity-AI", "Safety-Tech", "Finance-Tech", "Political-AI",
// 	"Energy-AI", "Manufacturing-AI", "Aerospace-AI", "Environmental-AI", "Public-Transport-AI",
// 	"Retail-Tech", "Sports-Training", "Legal-AI", "Metaverse-AI", "Hospitality-AI",
// 	"Medical-Research", "Digital-Marketing", "Scientific-Research", "Space-AI", "Weather-AI"}

// func GenerateUSer(num int) []*store.User {
// 	users := make([]*store.User, num)

// 	for i := 0; i < num; i++ {
// 		users[i] = &store.User{
// 			ID:       GenerateUUID(),
// 			Username: username[i%len(username)] + fmt.Sprintf("%d", i),
// 			Email:    username[i%len(username)] + fmt.Sprintf("%d", i) + "@example.com",
// 			Password: store.Password{},
// 		}
// 	}
// 	return users
// }

// func GeneratePosts(num int, users []*store.User) []*store.Post {
// 	posts := make([]*store.Post, num)
// 	user := users[rand.Intn(len(users))]
// 	for i := 0; i < num; i++ {
// 		posts[i] = &store.Post{
// 			ID:      GenerateUUID(),
// 			Title:   title[i%len(title)] + fmt.Sprintf("%d", i),
// 			Content: content[i%len(content)] + fmt.Sprintf("%d", i),
// 			Tags: []string{
// 				tags[i%len(tags)] + fmt.Sprintf("%d", i),
// 				tags[i%len(tags)] + fmt.Sprintf("%d", i),
// 				tags[i%len(tags)] + fmt.Sprintf("%d", i),
// 			},
// 			UserID: user.ID,
// 		}

// 	}
// 	return posts
// }

// func GenerateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
// 	comments := make([]*store.Comment, num)
// 	user := users[rand.Intn(len(users))]
// 	post := posts[rand.Intn(len(posts))]
// 	for i := 0; i < num; i++ {
// 		comments[i] = &store.Comment{
// 			ID:      GenerateUUID(),
// 			PostID:  post.ID,
// 			UserID:  user.ID,
// 			Comment: comment[i%len(comment)] + fmt.Sprintf("%d", i),
// 		}

// 	}
// 	return comments
// }
