# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.0.1 - 2024-11-06]
### Added
- Initial implementation of ShardedMap with support for generics.
- Set, Get, and Delete methods for managing key-value pairs.
- Count method to retrieve the total number of items in the map.
- Keys method to list all keys in the map.
- Clear method to remove all key-value pairs.
- Sharding functionality with a default of 32 shards.
- Atomic operations for thread-safe counting of items across shards.