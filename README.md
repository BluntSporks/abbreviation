# abbreviation
Reusable Golang library to provide functions to check the validity of abbreviations

## Purpose
To list all short names that I use in programs for reference.

## Status
Ready to use

## Rules
* Short names can't be ambiguous. However, if the names are different that no confusion can result, they are OK.
* Short names should have at least one vowel.
* Short names should usually end in a consonant, unless the vowel is needed for discrimination, for example, alg and algo.
* Short names should not have more than three consonants in a row.
* Short names should be pronounceable.
* Short names should be at least three letters long.
* A short name should be half as long as the original word or less, or just use the long form.
* All of the letters in the short form should be present in the long form and in the same order.
* Short should not be whole words that mean something else, for example, fact for factorial.
* Short names can be differentiated by ending, for example, fac and facl.
* Acronyms are OK as short names, but the acronym should already be well known. They are not listed here.
* Short names should be compatible with the Google Go language. They should be ideally already used in Go code.
* You can add s to a short name to form the plural, as long as it doesn't end in s, but do not add d to form past tense.
* Short names should be verified at [AllAcronyms.com](http://www.allacronyms.com/), unless it conflicts with above.
* There are a few exceptions to the above rules for common, well-established forms.

## List
* abbrev = abbreviate, abbreviation
* abs = absolute
* acc = accuracy, accurate
* acct = account
* accum = accumulate, accumulation, accumulator
* act = actual
* addr = address
* adn = addend
* agg = aggregate, aggregation
* alg = algebra
* algo = algorithm
* alloc = allocate, allocation
* anc = ancestor
* approx = approximate, approximation
* arch = architect, architecture
* arg = argument
* asc = ascending
* attr = attribute
* avg = average
* bin = binary
* binom = binomial
* bool = boolean
* buf = buffer
* cal = calendar
* calc = calculate, calculation, calculator, calculus
* cap = capitalize, capitalization, capacity
* ceil = ceiling
* cert = certify, certificate, certification
* chan = channel
* char = character
* coef = coefficient
* col = column
* coll = collect, collection
* comb = combine, combination
* comp = compare, comparison, compensate, compensation
* cond = condition
* config = configure, configuration
* conn = connect, connection, connector
* const = constant
* cont = continue, continuation, contain, container
* conv = convert, conversion
* cop = complex
* corr = correlate, correlation
* cos = cosine
* cot = cotangent
* covar = covariance
* cur = current
* dec = decimal, decrement, decode, decrypt
* decl = declare, declaration
* def = define, definition
* deg = degree
* del = delete, deletion
* delim = delimiter
* den = denominate, denomination, denominator
* dep = depend, dependent
* des = descendant
* desc = descending, describe, description
* dest = destine, destination
* det = determinant
* dev = deviate, deviation, develop, development
* diag = diagonal
* diam = diameter
* dict = dictionary
* diff = difference
* dim = dimension
* dir = directory, directive, direct, direction
* disc = discriminant
* dist = distribute, distribution, distributor, distance
* div = divide
* doc = document, documentation
* dom = domain
* dupe = duplicate, duplication
* elem = element
* empl = employee, employment
* enc = encode, encrypt
* ent = entropy
* enum = enumerate, enumeration, enumerator
* env = environment
* equiv = equivalent, equivalence
* err = error
* esc = escape
* est = estimate, estimation
* eval = evaluate, evaluation
* excl = exclude, exclusion
* excp = except, exception
* exec = execute, execution
* exp = exponent, expect
* expr = express, expression
* ext = extend, extension, external
* fac = factor
* facl = factorial
* fig = figure
* float = floating
* frac = fraction
* freq = frequent, frequency
* func = function
* gen = generate, generation, generator, general, generic
* geom = geometric
* guar = guarantee
* har = harmonic
* hex = hexadecimal
* horiz = horizontal
* ident = identification, identifier
* imag = imaginary
* img = image
* imm = immediate
* imp = implement, implementation
* in = input
* inc = increment
* incl = include, inclusion
* ind = indirect, indicate, indication
* indep = independent
* inf = infinite, infinity
* info = information
* init = initial, initialize, initialization
* ins = insert, insertion
* inst = instance, instantiate, instruction
* int = integer, internal
* intc = intercept
* intf = interface
* intx = intersection
* inv = inverse
* io = input/output
* iter = iterate, iteration, iterator
* kurt = kurtosis
* lang = language
* len = length
* lex = lexicon
* lib = library
* lim = limit, limitation
* lin = linear
* lit = literal
* loc = locate, location
* log = logarithm
* low = lowercase
* mach = machine
* mant = mantissa
* max = maximum
* med = medium, median
* mem = memory
* min = minimum, minute
* mod = modulo, modify, modification, module
* mul = multiply
* multi = multiple
* mut = mutate, mutation, mutable
* mutex = mutual exclusion
* mutl = mutual
* nat = natural
* neg = negate, negation, negative
* net = network
* num = number, numerate, numeration, numerator
* obj = object
* oct = octal
* op = operate, operation, operator
* opt = option
* orig = original, originate, origination
* out = output
* ovfl = overflow
* par = parent
* para = paragraph
* param = parameter
* parl = parallel
* pat = pattern
* perf = perform, performance
* perim = perimeter
* perm = permutate, permutation, permanent
* perp = perpendicular
* phon = phonetic
* pic = picture
* pos = positive, position
* poss = possible, possibly
* prec = precise, precision
* prep = prepare, preparation
* prev = previous
* pri = private, primary
* prio = priority, prioritize
* prob = probability
* proc = procedure, process
* prod = product
* prof = profile
* prog = program, progress
* proj = project
* pron = pronounce, pronunciation
* prop = proportion
* prot = protocol
* proto = prototype
* pub = public
* punct = punctuate, punctuation
* quad = quadratic
* quot = quotient
* rand = random
* rat = ratio, rational
* rec = recursive, record
* rect = rectangle
* recv = receive
* ref = reference
* reg = register, regular
* regexp = regular expression
* rel = relate, relation, relative
* rem = remainder
* rep = represent, representation, representative, repeat, repetition
* repl = replace, replacement
* repo = repository
* req = require, requirement
* res = resource, result
* resp = respond, response, responsive
* ret = return
* rev = reverse
* san = sanitize, sanitization
* sched = schedule
* sci = science, scientific
* sec = second, secondary, secure, security, secant
* sect = section
* seg = segment
* sel = select, selection
* sent = sentence
* sep = separate, separation, separator
* seq = sequence
* sig = signal, signature
* sin = sine
* sock = socket
* spec = specify, specification, specific
* stat = statistic
* struct = structure
* sub = subtract
* surf = surface
* syl = syllable
* sym = symbol
* symm = symmetry
* sync = synchronized
* sys = system
* tan = tangent
* tech = technology
* temp = template, temporary
* term = terminate, termination, terminal
* tok = token, tokenize, tokenization
* trig = trigonometry
* up = uppercase
* upd = update
* util = utility
* val = value, validate, validation
* var = variance, variable
* vec = vector
* ver = version, verify, verification
* vert = vertical
* vis = visible, visibility
* vol = volume
* win = window
* xfer = transfer
* xlat = translate, translation
* xmis = transmission
* xmit = transmit
* xor = exclusive or

## Not in combination
These are OK to use as whole words, or at the beginning or end of a camelcased word, but try not to use them in the
middle of words due to their lack of vowels.
* cmd = command
* cmp = compare, comparison
* cnt = count
* csc = cosecant
* ctl = control
* ctx = context
* cvt = convert
* dst = destination
* fld = field
* fmt = format
* fwd = forward
* hdl = handle
* hdr = header
* i18n = internationalization
* l10n = localization
* msg = message
* pkg = package
* ptr = pointer
* qry = query
* src = source
* srv = server
* std = standard
* stm = statement
* str = string
* tbl = table
* tmp = temporary

## Not abbreviated
Here is a list of words that I do not abbreviate, even though some of them already have common abbreviations like line.
* array
* default
* empty
* equal
* file
* index
* line
* name
* test
* text
